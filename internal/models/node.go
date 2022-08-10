package models

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"github.com/samredai/sludge/internal/api"
	"gopkg.in/yaml.v3"
)

// Get information on all nodes on the SLURM cluster and place it in a collection
func GetAllNodes() *Collection[api.NodeInfo] {
	slurm_response, err := api.RpcGetAllNodes()
	if err != nil {
		log.Fatal(err)
	}
	nodes, err := api.ParseNodeInfoMsg(slurm_response)
	if err != nil {
		log.Fatal(err)
	}

	rowLength := reflect.TypeOf(api.NodeInfo{}).NumField()
	var rows []Row[api.NodeInfo]
	for _, node := range nodes {
		rows = append(rows, Row[api.NodeInfo]{
			Data:   node,
			Length: rowLength,
		})
	}

	c := Collection[api.NodeInfo]{
		Items: rows,
		Title: "Nodes",
		Count: len(rows),
		Headers: func(r Row[api.NodeInfo]) []string {
			var header []string
			structValue := r.Data
			fields := reflect.TypeOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				header = append(header, fields.Field(i).Name)
			}
			return header
		},
		Row: func(r Row[api.NodeInfo]) []string {
			var vals []string

			structValue := r.Data
			fields := reflect.TypeOf(structValue)
			values := reflect.ValueOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				value := values.Field(i)

				switch value.Kind() {
				case reflect.String:
					vals = append(vals, value.String())
				case reflect.Int:
					vals = append(vals, strconv.Itoa(int(value.Int())))
				default:
					vals = append(vals, "")
				}
			}
			return vals
		},
		RowData: func(c Collection[api.NodeInfo]) []float64 {
			var vals []float64

			rand.Seed(time.Now().UTC().UnixNano())
			for range c.Items {
				datum := 0 + rand.Float64()*(100-0)
				vals = append(vals, datum)
			}
			return vals
		},
	}
	return &c
}

// Get info on a single node in the SLURM cluster
func GetNodeInfo(name string) string {
	slurm_response, err := api.RpcGetNode(name)
	if err != nil {
		log.Fatal(err)
	}
	nodes, err := api.ParseNodeInfoMsg(slurm_response)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range nodes {
		if n.Name == name {

			data, err := yaml.Marshal(&n)

			if err != nil {
				log.Fatal(err)
			}

			yamlStr := string(data)
			return yamlStr
		}
	}
	return fmt.Sprintf("Could not find node info for name %d", name)
}
