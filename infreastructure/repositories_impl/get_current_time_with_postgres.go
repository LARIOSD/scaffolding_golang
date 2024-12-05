package repositories_impl

import (
	"context"
	"genesis/pos/domain/ports/infrastructure"
	"log"
	"time"
)

type GetCurrentTimeWithPostgresImpl struct {
	DB infrastructure.DatabaseConnectionPgxInterface
}

func (getTime *GetCurrentTimeWithPostgresImpl) ExecuteImpl() (time.Time, error) {
	type timeStruct struct {
		CurrentTime time.Time `json:"CurrentTime"`
	}

	timeInstance := &timeStruct{}
	query := `SELECT NOW() AS "CurrentTime";`

	conn := getTime.DB.GetDatabaseConnectionPgx().UsePgxDatabaseConnectionDriver()

	resultCurrentTime := conn.QueryRow(context.Background(), query)

	err := resultCurrentTime.Scan(&timeInstance.CurrentTime)

	if err != nil {
		log.Println("Error ::: ", err)
	}

	return timeInstance.CurrentTime, nil
}
