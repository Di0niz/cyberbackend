package config

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlMaster1 = flag.String("mysqlmaster1",
		"root:example@tcp(localhost:3306)/cybergame?&charset=utf8&interpolateParams=true",
		"mySql1 Master addr")
	mysqlSlave1 = flag.String("mysqlslave1",
		"root:example@tcp(localhost:3406)/cybergame?&charset=utf8&interpolateParams=true",
		"mySql1 Master addr")

	mysqlMaster2 = flag.String("mysqlmaster2",
		"root:example@tcp(localhost:3307)/cybergame?&charset=utf8&interpolateParams=true",
		"mySql2 Master addr")
	mysqlSlave2 = flag.String("mysqlslave2",
		"root:example@tcp(localhost:3407)/cybergame?&charset=utf8&interpolateParams=true",
		"mySql2 Slave addr")
)

type (

	// определяем список баз данных для подключения

	DBPool struct {
		Connections    []*DBConnection
		TimeCheckAlive time.Duration
	}

	// параметр подключения к базе данных
	DBConnection struct {
		Master       *sql.DB
		Slave        *sql.DB
		Number       int
		IsLiveMaster bool
		IsLiveSlave  bool
	}

	// определяем интерфейс работы с таблицами
	DBTable interface {
		Create(db *sql.DB) error
		Update(db *sql.DB) error
		//Custom(db *sql.DB, sqlstring string) error
		ShardNumber() int
	}

	// определяем интерфейс работы с таблицами
	DBList interface {
		List(db *sql.DB) ([]*DBTable, error)
	}
)

func (con *DBConnection) String() string {

	return fmt.Sprintf("{Number: %d, IsLive: %t}", con.Number, con.IsLiveMaster && con.IsLiveSlave)
}

func PrepareDatabase(dns string) *sql.DB {
	db, err := sql.Open("mysql", dns)

	if err != nil {
		db.SetMaxOpenConns(20)
	}

	return db

}

func (pool *DBPool) CheckLiveDatabase() {

	go func(pool *DBPool) {
		for {
			for _, con := range pool.Connections {

				con.IsLiveMaster = con.Master.Ping() == nil
				con.IsLiveSlave = con.Slave.Ping() == nil

				//fmt.Println("con.Master:", con.IsLiveMaster, con)
				//fmt.Println("con.Slave:", con.IsLiveSlave, con)

			}
			time.Sleep(time.Second * pool.TimeCheckAlive)
		}
	}(pool)

}

// создаем шардировнную базу данных
func MySQLConnect() *DBPool {

	pool := &DBPool{
		Connections: []*DBConnection{
			&DBConnection{
				Master: PrepareDatabase(*mysqlMaster1),
				Slave:  PrepareDatabase(*mysqlSlave1),
				Number: 1,
			},
			&DBConnection{
				Master: PrepareDatabase(*mysqlMaster2),
				Slave:  PrepareDatabase(*mysqlSlave2),
				Number: 2,
			},
		},
		TimeCheckAlive: 3,
	}

	pool.CheckLiveDatabase()

	return pool
}

func (pool *DBPool) Create(item DBTable) error {

	sharNumber := item.ShardNumber()

	// запись делам только в мастер

	con := pool.Connections[sharNumber]

	if con.IsLiveMaster {

		return item.Create(con.Master)

	} else {
		return fmt.Errorf("We have some problem with Server! (((")
	}

	return nil

}

func (pool *DBPool) Update(item DBTable) error {

	sharNumber := item.ShardNumber()

	// запись делам только в мастер

	con := pool.Connections[sharNumber]

	if con.IsLiveMaster {

		return item.Update(con.Master)

	} else {
		return fmt.Errorf("We have some problem with Server! (((")
	}

	return nil

}

// Логика работы списков
// объединяем все результаты в один slice
// потом его будем фильтровать и сортировать
func (pool *DBPool) List(L DBList) error {

	for con := range pool.Connections {

		// по умолчанию читаем со slave
		if con.IsLiveSlave {

			L.List(con)

			// если slave не доступен, читаем из мастера
		} else if con.IsLiveMaster {
			L.List(con)

		}

	}

	return nil
}
