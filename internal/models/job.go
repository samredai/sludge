package models

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"github.com/samredai/sludge/internal/api"
	"gopkg.in/yaml.v3"
)

func GetJobs(ctx context.Context) *Collection[api.JobDetails] {
	return getAllJobs(ctx)
}

func getAllJobs(ctx context.Context) *Collection[api.JobDetails] {
	username, _ := ctx.Value(USERNAME_KEY).(string)
	account, _ := ctx.Value(ACCOUNT_KEY).(string)
	slurm_response, err := api.RpcGetAllJobs()
	if err != nil {
		log.Fatal(err)
	}
	jobs, err := api.ParseJobInfoMsg(slurm_response)
	if err != nil {
		log.Fatal(err)
	}

	rowLength := reflect.TypeOf(api.JobDetails{}).NumField()
	var rows []Row[api.JobDetails]
	if username == "" && account == "" {
		for _, job := range jobs { // Append all jobs
			rows = append(rows, Row[api.JobDetails]{
				Data:   job,
				Length: rowLength,
			})
		}
	} else {
		for _, job := range jobs {
			if (username == "" || username == job.Username) && (account == "" || account == job.Account) { // Only append if matches username/account filter
				rows = append(rows, Row[api.JobDetails]{
					Data:   job,
					Length: rowLength,
				})
			}
		}
	}

	c := Collection[api.JobDetails]{
		Items: rows,
		Title: "Jobs",
		Count: len(rows),
		Headers: func(r Row[api.JobDetails]) []string {
			var header []string
			structValue := r.Data
			fields := reflect.TypeOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				header = append(header, fields.Field(i).Name)
			}
			return header
		},
		Row: func(r Row[api.JobDetails]) []string {
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
		RowData: func(c Collection[api.JobDetails]) []float64 {
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

func getAllUsersJobs(user int) *Collection[api.JobDetails] {
	slurm_response, err := api.RpcGetAllUsersJobs(user)
	if err != nil {
		log.Fatal(err)
	}
	jobs, err := api.ParseJobInfoMsg(slurm_response)
	if err != nil {
		log.Fatal(err)
	}

	rowLength := reflect.TypeOf(api.JobDetails{}).NumField()
	var rows []Row[api.JobDetails]
	for _, job := range jobs {
		rows = append(rows, Row[api.JobDetails]{
			Data:   job,
			Length: rowLength,
		})
	}

	c := Collection[api.JobDetails]{
		Items: rows,
		Title: "Jobs",
		Count: len(rows),
		Headers: func(r Row[api.JobDetails]) []string {
			var header []string
			structValue := r.Data
			fields := reflect.TypeOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				header = append(header, fields.Field(i).Name)
			}
			return header
		},
		Row: func(r Row[api.JobDetails]) []string {
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
		RowData: func(c Collection[api.JobDetails]) []float64 {
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

// Get info on a single active job
func GetJobInfo(id int) string {
	slurm_response, err := api.RpcGetJob(id)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	jobs, err := api.ParseJobInfoMsg(slurm_response)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	for _, job := range jobs {
		if job.ID == id {

			data, err := yaml.Marshal(&job)

			if err != nil {
				log.Fatal(err)
			}

			yamlStr := string(data)
			return yamlStr
		}
	}
	return fmt.Sprintf("Could not find job info for ID %d", id)
}
