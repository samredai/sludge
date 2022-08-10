package models

import (
	"context"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"github.com/samredai/sludge/internal/api"
)

type ClusterUtilizationData struct {
	Cluster   string
	Allocated string
	Down      string
	PLNDDown  string
	Idle      string
	Planned   string
	Reported  string
}

// Run the "Cluster Utilization" report and place the results in a collection
func GetClusterUtilizationReport(ctx context.Context) *Collection[ClusterUtilizationData] {
	records, _ := api.GetParsableReport("Cluster", "Utilization")
	var rows []Row[ClusterUtilizationData]
	for _, record := range records {
		rows = append(rows, Row[ClusterUtilizationData]{
			Data: ClusterUtilizationData{
				Cluster:   record[0],
				Allocated: record[1],
				Down:      record[2],
				PLNDDown:  record[3],
				Idle:      record[4],
				Planned:   record[5],
				Reported:  record[6],
			},
			Length: 7,
		})
	}

	c := Collection[ClusterUtilizationData]{
		Items: rows,
		Title: "Cluster Utilization",
		Count: len(rows),
		Headers: func(r Row[ClusterUtilizationData]) []string {
			var header []string
			structValue := r.Data
			fields := reflect.TypeOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				header = append(header, fields.Field(i).Name)
			}
			return header
		},
		Row: func(r Row[ClusterUtilizationData]) []string {
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
		RowData: func(c Collection[ClusterUtilizationData]) []float64 {
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

type ClusterAccountUtilizationByUserData struct {
	Cluster    string
	Account    string
	Login      string
	ProperName string
	Used       string
	Energy     string
}

// Run the "Cluster AccountUtilizationByUser" report and place the results in a collection
func GetClusterAccountUtilizationByUserReport(ctx context.Context) *Collection[ClusterAccountUtilizationByUserData] {
	username, _ := ctx.Value(USERNAME_KEY).(string)
	account, _ := ctx.Value(ACCOUNT_KEY).(string)
	records, _ := api.GetParsableReport("Cluster", "AccountUtilizationByUser")
	var rows []Row[ClusterAccountUtilizationByUserData]
	if username == "" && account == "" {
		for _, record := range records {
			rows = append(rows, Row[ClusterAccountUtilizationByUserData]{
				Data: ClusterAccountUtilizationByUserData{
					Cluster:    record[0],
					Account:    record[1],
					Login:      record[2],
					ProperName: record[3],
					Used:       record[4],
					Energy:     record[5],
				},
				Length: 6,
			})
		}
	} else {
		for _, record := range records {
			if (username == "" || username == record[2]) && (account == "" || account == record[1]) {
				rows = append(rows, Row[ClusterAccountUtilizationByUserData]{
					Data: ClusterAccountUtilizationByUserData{
						Cluster:    record[0],
						Account:    record[1],
						Login:      record[2],
						ProperName: record[3],
						Used:       record[4],
						Energy:     record[5],
					},
					Length: 6,
				})
			}
		}
	}

	c := Collection[ClusterAccountUtilizationByUserData]{
		Items: rows,
		Title: "Cluster Account Utilization By User",
		Count: len(rows),
		Headers: func(r Row[ClusterAccountUtilizationByUserData]) []string {
			var header []string
			structValue := r.Data
			fields := reflect.TypeOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				header = append(header, fields.Field(i).Name)
			}
			return header
		},
		Row: func(r Row[ClusterAccountUtilizationByUserData]) []string {
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
		RowData: func(c Collection[ClusterAccountUtilizationByUserData]) []float64 {
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

type ClusterUserUtilizationByAccountData struct {
	Cluster    string
	Login      string
	ProperName string
	Account    string
	Used       string
	Energy     string
}

// Run the "Cluster UserUtilizationByAccount" report and place the results in a collection
func GetClusterUserUtilizationByAccountReport(ctx context.Context) *Collection[ClusterUtilizationData] {
	records, _ := api.GetParsableReport("Cluster", "UserUtilizationByAccount")
	var rows []Row[ClusterUtilizationData]
	for _, record := range records {
		rows = append(rows, Row[ClusterUtilizationData]{
			Data: ClusterUtilizationData{
				Cluster:   record[0],
				Allocated: record[1],
				Down:      record[2],
				PLNDDown:  record[3],
				Idle:      record[4],
				Planned:   record[5],
				Reported:  record[6],
			},
			Length: 7,
		})

	}

	c := Collection[ClusterUtilizationData]{
		Items: rows,
		Title: "Cluster User Utilization By Account",
		Count: len(rows),
		Headers: func(r Row[ClusterUtilizationData]) []string {
			var header []string
			structValue := r.Data
			fields := reflect.TypeOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				header = append(header, fields.Field(i).Name)
			}
			return header
		},
		Row: func(r Row[ClusterUtilizationData]) []string {
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
		RowData: func(c Collection[ClusterUtilizationData]) []float64 {
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

type JobSizesByAccountData struct {
	Cluster          string
	Account          string
	CPUs_0to49       string
	CPUs_50to249     string
	CPUs_250to499    string
	CPUs_500to999    string
	CPUs_Over1000    string
	PercentOfCluster string
}

// Run the "Job SizesByAccount" report and place the results in a collection
func GetJobSizesByAccountReport(ctx context.Context) *Collection[ClusterUtilizationData] {
	records, _ := api.GetParsableReport("Job", "SizesByAccount")
	var rows []Row[ClusterUtilizationData]
	for _, record := range records {
		rows = append(rows, Row[ClusterUtilizationData]{
			Data: ClusterUtilizationData{
				Cluster:   record[0],
				Allocated: record[1],
				Down:      record[2],
				PLNDDown:  record[3],
				Idle:      record[4],
				Planned:   record[5],
				Reported:  record[6],
			},
			Length: 7,
		})
	}

	c := Collection[ClusterUtilizationData]{
		Items: rows,
		Title: "Job Sizes By Account",
		Count: len(rows),
		Headers: func(r Row[ClusterUtilizationData]) []string {
			var header []string
			structValue := r.Data
			fields := reflect.TypeOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				header = append(header, fields.Field(i).Name)
			}
			return header
		},
		Row: func(r Row[ClusterUtilizationData]) []string {
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
		RowData: func(c Collection[ClusterUtilizationData]) []float64 {
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

type ReservationUtilizationData struct {
	Cluster   string
	Name      string
	Start     string
	End       string
	TRESName  string
	Allocated string
	Idle      string
}

// Run the "Reservation Utilization" report and place the results in a collection
func GetReservationUtilizationReport(ctx context.Context) *Collection[ClusterUtilizationData] {
	records, _ := api.GetParsableReport("Reservation", "Utilization")
	var rows []Row[ClusterUtilizationData]
	for _, record := range records {
		rows = append(rows, Row[ClusterUtilizationData]{
			Data: ClusterUtilizationData{
				Cluster:   record[0],
				Allocated: record[1],
				Down:      record[2],
				PLNDDown:  record[3],
				Idle:      record[4],
				Planned:   record[5],
				Reported:  record[6],
			},
			Length: 7,
		})
	}

	c := Collection[ClusterUtilizationData]{
		Items: rows,
		Title: "Reservation Utilization",
		Count: len(rows),
		Headers: func(r Row[ClusterUtilizationData]) []string {
			var header []string
			structValue := r.Data
			fields := reflect.TypeOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				header = append(header, fields.Field(i).Name)
			}
			return header
		},
		Row: func(r Row[ClusterUtilizationData]) []string {
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
		RowData: func(c Collection[ClusterUtilizationData]) []float64 {
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

type UserTopUsageData struct {
	Cluster    string
	Login      string
	ProperName string
	Account    string
	Used       string
	Energy     string
}

// Run the "User TopUsage" report and place the results in a collection
func GetUserTopUsageReport(ctx context.Context) *Collection[ClusterUtilizationData] {
	records, _ := api.GetParsableReport("User", "TopUsage")
	var rows []Row[ClusterUtilizationData]
	for _, record := range records {
		rows = append(rows, Row[ClusterUtilizationData]{
			Data: ClusterUtilizationData{
				Cluster:   record[0],
				Allocated: record[1],
				Down:      record[2],
				PLNDDown:  record[3],
				Idle:      record[4],
				Planned:   record[5],
				Reported:  record[6],
			},
			Length: 7,
		})
	}

	c := Collection[ClusterUtilizationData]{
		Items: rows,
		Title: "User Top Usage",
		Count: len(rows),
		Headers: func(r Row[ClusterUtilizationData]) []string {
			var header []string
			structValue := r.Data
			fields := reflect.TypeOf(structValue)

			num := fields.NumField()

			for i := 0; i < num; i++ {
				header = append(header, fields.Field(i).Name)
			}
			return header
		},
		Row: func(r Row[ClusterUtilizationData]) []string {
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
		RowData: func(c Collection[ClusterUtilizationData]) []float64 {
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
