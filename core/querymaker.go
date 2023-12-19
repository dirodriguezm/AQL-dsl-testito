package core

import (
	"encoding/json"
	"fmt"
	"poc/aql/grammar"
)

type QueryArguments struct {
	Oid string `json:"oid"`
}

type UseCase struct{}

func (uc *UseCase) getMethod(entity string) func(queryArgs *QueryArguments) interface{} {
	switch entity {
	case "Detection":
		lcUseCase := LightcurveUseCase{}
		return func(queryArgs *QueryArguments) interface{} {
			return lcUseCase.GetDetection(queryArgs)
		}
	case "NonDetection":
		lcUseCase := LightcurveUseCase{}
		return func(queryArgs *QueryArguments) interface{} {
			return lcUseCase.GetNonDetection(queryArgs)
		}
	case "Object":
		objUseCase := ObjectUseCase{}
		return func(queryArgs *QueryArguments) interface{} {
			return objUseCase.GetObject(queryArgs)
		}
	default:
		panic(fmt.Sprintf("entity %s not supported", entity))
	}
}

func getQueryArguments(query *grammar.Query) *QueryArguments {
	queryArguments := QueryArguments{}
	for _, filter := range query.Filters {
		switch filter.Field {
		case "oid":
			queryArguments.Oid = filter.Value.(grammar.String).String
		default:
			fmt.Println(fmt.Sprintf("Filter \"%s\" not supported. Ignoring it.", filter.Field))
		}
	}
	return &queryArguments
}

func MakeQuery(query *grammar.Query) ([]byte, error) {
	// this method should determine what kind of query is being made
	// it could be getDetection, getNonDetection, getLightCurve
	// and should determine what arguments should be used
	fmt.Printf("query is %+v\n", query)
	useCase := UseCase{}
	method := useCase.getMethod(query.Entity)
	queryArguments := getQueryArguments(query)
	result := method(queryArguments)
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("error marshaling result to JSON: %v", err)
	}
	return jsonResult, nil
}
