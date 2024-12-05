package container

import (
	"genesis/pos/application/service_application"
	use_case2 "genesis/pos/application/use_case"
	service_application2 "genesis/pos/domain/ports/application/service_application"
	"genesis/pos/domain/ports/application/use_case"
	"genesis/pos/domain/ports/infrastructure"
	"genesis/pos/infreastructure/db/mongo/drivers"
	drivers2 "genesis/pos/infreastructure/db/postgres/drivers"
	"genesis/pos/infreastructure/external_http"
	"genesis/pos/infreastructure/external_http/net_http"
	"genesis/pos/infreastructure/repositories_impl"
	"os"
)

var DatabaseConnectionToMongo infrastructure.DatabaseConnectionMongoInterface

func ResolveDatabaseConnectionToMongo() infrastructure.DatabaseConnectionMongoInterface {
	if DatabaseConnectionToMongo != nil {
		return DatabaseConnectionToMongo
	} else {
		return &drivers.ConfigConnectionMongo{
			UrlConnection: os.Getenv("CONNECTION_MONGOOSE"),
		}
	}

}

var DatabaseConnectionToLecWithPgx infrastructure.DatabaseConnectionPgxInterface

func ResolveDatabaseConnectionToLecWithPgx() infrastructure.DatabaseConnectionPgxInterface {
	if DatabaseConnectionToLecWithPgx != nil {
		return DatabaseConnectionToLecWithPgx
	}
	return &drivers2.ConfigConnectionPgx{
		UrlToConnect: os.Getenv("CONNECTION_POSTGRESQL"),
	}

}

var ClientHttp external_http.ClientHTTPInterface

func ResolveClientHttp() external_http.ClientHTTPInterface {
	if ClientHttp != nil {
		return ClientHttp
	}
	return &net_http.NetHTTPClient{}
}

type ApplicationParameters struct {
	HostServer string
}

var AppParams ApplicationParameters

var ChargeInitialParamsWithPostgres *repositories_impl.ChargeInitialParamsWithPostgres

func ResolveChargeInitialParamsWithPostgres() *repositories_impl.ChargeInitialParamsWithPostgres {
	if ChargeInitialParamsWithPostgres != nil {
		return ChargeInitialParamsWithPostgres
	}
	return &repositories_impl.ChargeInitialParamsWithPostgres{
		DB: DatabaseConnectionToLecWithPgx,
	}
}

var GetEnvironmentRepository infrastructure.GetEnvironmentRepository

func ResolveGetEnvironmentRepository() infrastructure.GetEnvironmentRepository {
	if GetEnvironmentRepository != nil {
		return GetEnvironmentRepository
	}
	return &repositories_impl.GetEnvironmentImpl{}
}

var GetCurrentTimeRepository infrastructure.GetCurrentTimeRepository

func ResolveGetCurrentTimeRepository() infrastructure.GetCurrentTimeRepository {
	if GetCurrentTimeRepository != nil {
		return GetCurrentTimeRepository
	}

	return &repositories_impl.GetCurrentTimeWithPostgresImpl{
		DB: DatabaseConnectionToLecWithPgx,
	}
}

var GetCurrentTimeUseCase use_case.GetCurrentTimeUseCase

func ResolveGetCurrentTimeUseCase() use_case.GetCurrentTimeUseCase {
	if GetCurrentTimeUseCase != nil {
		return GetCurrentTimeUseCase
	}
	return &use_case2.GetCurrentTimeUseCaseImpl{
		GetCurrentTimeRepo: GetCurrentTimeRepository,
	}
}

var GetSaleRepository infrastructure.GetSalesRepository

func ResolveGetSaleRepository() infrastructure.GetSalesRepository {
	if GetSaleRepository != nil {
		return GetSaleRepository
	}
	return &repositories_impl.GetSalesWithMongoImpl{
		MongoDB: DatabaseConnectionToMongo,
	}
}

var GetSaleUseCase use_case.GetSalesUseCase

func ResolveGetSaleUseCase() use_case.GetSalesUseCase {
	if GetSaleUseCase != nil {
		return GetSaleUseCase
	}
	return &use_case2.GetSalesUseCaseImpl{
		GetSalesUseCaseRepo: GetSaleRepository,
	}
}

var GetCurrentTimeService service_application2.GetCurrentTimeService

func ResolveGetCurrentTimeService() service_application2.GetCurrentTimeService {
	if GetCurrentTimeService != nil {
		return GetCurrentTimeService
	}

	return &service_application.GetCurrentTimeServiceImpl{
		GetCurrentTimeServiceUseCase: GetCurrentTimeUseCase,
		GetSalesUseCase:              GetSaleUseCase,
	}
}

func StartContainer() {
	DatabaseConnectionToMongo = ResolveDatabaseConnectionToMongo()
	DatabaseConnectionToLecWithPgx = ResolveDatabaseConnectionToLecWithPgx()
	ClientHttp = ResolveClientHttp()

	//Parameters of application
	ChargeInitialParamsWithPostgres = ResolveChargeInitialParamsWithPostgres()
	parameters, err := ChargeInitialParamsWithPostgres.ChargeInitialParameters()
	if err != nil {
		return
	}
	AppParams.HostServer = parameters

	//Repositories general
	GetEnvironmentRepository = ResolveGetEnvironmentRepository()

	//Repositories
	GetCurrentTimeRepository = ResolveGetCurrentTimeRepository()
	GetSaleRepository = ResolveGetSaleRepository()

	//Uses cases
	GetCurrentTimeUseCase = ResolveGetCurrentTimeUseCase()
	GetSaleUseCase = ResolveGetSaleUseCase()

	//Services
	GetCurrentTimeService = ResolveGetCurrentTimeService()
}
