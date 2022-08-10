package api

/*
#cgo pkg-config: slurm
#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include "slurm/slurm.h"
#include "slurm/slurm_errno.h"
*/
import "C"

import (
	"fmt"
	"log"
	"time"
	"unsafe"
)

// Make an RPC call to the SLURM cluster to get information on all jobs
// Make sure to free the response with C.slurm_free_job_info_msg(slurm_response)
func RpcGetAllJobs() (*C.struct_job_info_msg, error) {
	var slurm_response *C.job_info_msg_t
	exit_reason := C.slurm_load_jobs(C.time_t(0), &slurm_response, C.uint16_t(1))
	if exit_reason != 0 {
		log.Fatal(exit_reason)
		C.slurm_free_job_info_msg(slurm_response)
		return slurm_response, fmt.Errorf("Error loading jobs")
	}
	return slurm_response, nil
}

// Make an RPC call to the SLURM cluster to get information on all jobs for a specific user
// Make sure to free the response with C.slurm_free_job_info_msg(slurm_response)
func RpcGetAllUsersJobs(user int) (*C.struct_job_info_msg, error) {
	var slurm_response *C.job_info_msg_t
	exit_reason := C.slurm_load_job_user(&slurm_response, C.uint(user), C.uint16_t(1))
	if exit_reason != 0 {
		log.Fatal(exit_reason)
		C.slurm_free_job_info_msg(slurm_response)
		return slurm_response, fmt.Errorf("Error loading jobs")
	}
	return slurm_response, nil
}

// Make an RPC call to the SLURM cluster to get information on a specific job
// Make sure to free the response with C.slurm_free_job_info_msg(slurm_response)
func RpcGetJob(id int) (*C.struct_job_info_msg, error) {
	var slurm_response *C.job_info_msg_t
	exit_reason := C.slurm_load_job(&slurm_response, C.uint(id), C.uint16_t(1))
	if exit_reason != 0 {
		log.Fatal(exit_reason)
		C.slurm_free_job_info_msg(slurm_response)
		return slurm_response, fmt.Errorf("Error loading jobs")
	}
	return slurm_response, nil
}

// Parse the response of an RPC call to the SLURM cluster to receive information on a specific job
func ParseJobInfoMsg(slurm_response *C.struct_job_info_msg) ([]JobDetails, error) {
	var jobs_info []JobDetails
	cJobArray := slurm_response.job_array
	length := int(slurm_response.record_count)
	jobs := unsafe.Slice(cJobArray, length)

	for _, job := range jobs {
		info := JobDetails{
			ID:                int(job.job_id),
			Account:           C.GoString(job.account),
			Username:          C.GoString(job.user_name),
			UserID:            int(job.user_id),
			WorkDir:           C.GoString(job.work_dir),
			JobName:           C.GoString(job.name),
			SubmitTime:        time.Unix(int64(job.submit_time), 0),
			StartTime:         time.Unix(int64(job.start_time), 0),
			LastModTime:       time.Unix(int64(job.suspend_time), 0),
			SystemComment:     C.GoString(job.system_comment),
			TimeLimit:         int(job.time_limit),
			TimeMin:           int(job.time_min),
			ThreadsPerCore:    int(job.threads_per_core),
			StateReason:       int(job.state_reason),
			StateDescription:  C.GoString(job.state_desc),
			StdErr:            C.GoString(job.std_err),
			StdIn:             C.GoString(job.std_in),
			StdOut:            C.GoString(job.std_out),
			Shared:            int(job.shared),
			ShowFlags:         int(job.show_flags),
			ScheduledNodes:    C.GoString(job.sched_nodes),
			ReservationName:   C.GoString(job.resv_name),
			RequiredNodes:     C.GoString(job.req_nodes),
			Priority:          int(job.priority),
			PreSusTime:        time.Unix(int64(job.pre_sus_time), 0),
			MinTmpDiskPerNode: int(job.pn_min_tmp_disk),
			MinCPUsPerNode:    int(job.pn_min_cpus),
			MinMemPerNode:     int(job.pn_min_memory),
			Prefer:            C.GoString(job.prefer),
			Partition:         C.GoString(job.partition),
			NumTasks:          int(job.num_tasks),
			NumNodes:          int(job.num_nodes),
			NumCPUs:           int(job.num_cpus),
			NumTasksPerCore:   int(job.ntasks_per_core),
			NumTasksPerTres:   int(job.ntasks_per_tres),
			NumTasksPerNode:   int(job.ntasks_per_node),
			NumTasksPerSocket: int(job.ntasks_per_socket),
			NumTasksPerBoard:  int(job.ntasks_per_board),
			RestartCount:      int(job.restart_cnt),
			Nodes:             C.GoString(job.nodes),
			Network:           C.GoString(job.network),
			MemPerTres:        C.GoString(job.mem_per_tres),
			MaxNodes:          int(job.max_nodes),
			MaxCPUs:           int(job.max_cpus),
			MailUser:          C.GoString(job.mail_user),
			MailType:          int(job.mail_type),
			Licenses:          C.GoString(job.licenses),
			LastSchedEval:     time.Unix(int64(job.last_sched_eval), 0),
			JobState:          int(job.job_state),
			HetJobOffset:      int(job.het_job_offset),
			HetJobIdSet:       C.GoString(job.het_job_id_set),
			HetJobID:          int(job.het_job_id),
			GroupID:           int(job.group_id),
			GresTotal:         C.GoString(job.gres_total),
			GresDetailCnt:     int(job.gres_detail_cnt),
			FedOriginStr:      C.GoString(job.fed_origin_str),
			Features:          C.GoString(job.features),
			ExitCode:          int(job.exit_code),
			ExcludedNodes:     C.GoString(job.exc_nodes),
			EndTime:           time.Unix(int64(job.end_time), 0),
			EligibleTime:      time.Unix(int64(job.eligible_time), 0),
			DerivedExitCode:   int(job.derived_ec),
			Dependency:        C.GoString(job.dependency),
			DelayBoot:         int(job.delay_boot),
			Deadline:          time.Unix(int64(job.deadline), 0),
			Cronspec:          C.GoString(job.cronspec),
			CpusPerTres:       C.GoString(job.cpus_per_tres),
			CpuFreqGov:        int(job.cpu_freq_gov),
			CpuFreqMax:        int(job.cpu_freq_max),
			CpuFreqMin:        int(job.cpu_freq_min),
			CpusPerTask:       int(job.cpus_per_task),
			CoresPerSocket:    int(job.cores_per_socket),
			CoreSpec:          int(job.core_spec),
			Contiguous:        int(job.contiguous),
			Container:         C.GoString(job.container),
			Comment:           C.GoString(job.comment),
			Command:           C.GoString(job.command),
			ClusterFeatures:   C.GoString(job.cluster_features),
			Cluster:           C.GoString(job.cluster),
			BurstBufferState:  C.GoString(job.burst_buffer_state),
			BurstBuffer:       C.GoString(job.burst_buffer),
			BoardsPerNode:     int(job.boards_per_node),
			BitFlags:          int(job.bitflags),
			BatchHost:         C.GoString(job.batch_host),
			BatchFlag:         int(job.batch_flag),
			BatchFeatures:     C.GoString(job.batch_features),
			AssocId:           int(job.assoc_id),
			ArrayTaskStr:      C.GoString(job.array_task_str),
			ArrayMaxTasks:     int(job.array_max_tasks),
			ArrayTaskId:       int(job.array_task_id),
			ArrayJobId:        int(job.array_job_id),
			AllocSid:          int(job.alloc_sid),
			AllocNode:         C.GoString(job.alloc_node),
			AdminComment:      C.GoString(job.admin_comment),
			AccrueTime:        time.Unix(int64(job.accrue_time), 0),
		}
		jobs_info = append(jobs_info, info)
	}
	C.slurm_free_job_info_msg(slurm_response)
	return jobs_info, nil
}

// Make an RPC call to the SLURM cluster to get information on all nodes
// Make sure to free the response with C.slurm_free_node_info_msg(slurm_response)
func RpcGetAllNodes() (*C.struct_node_info_msg, error) {
	var slurm_response *C.node_info_msg_t
	exit_reason := C.slurm_load_node(C.time_t(0), &slurm_response, C.uint16_t(1))
	if exit_reason != 0 {
		log.Fatal(exit_reason)
		C.slurm_free_node_info_msg(slurm_response)
		return slurm_response, fmt.Errorf("Error loading nodes")
	}
	return slurm_response, nil
}

// Make an RPC call to the SLURM cluster to get information on a sepcific node
// Make sure to free the response with C.slurm_free_node_info_msg(slurm_response)
func RpcGetNode(name string) (*C.struct_node_info_msg, error) {
	var slurm_response *C.node_info_msg_t
	exit_reason := C.slurm_load_node_single(&slurm_response, C.CString(name), C.uint16_t(1))
	if exit_reason != 0 {
		log.Fatal(exit_reason)
		C.slurm_free_node_info_msg(slurm_response)
		return slurm_response, fmt.Errorf("Error loading node %s", name)
	}
	return slurm_response, nil
}

// Parse the response of an RPC call to the SLURM cluster to receive information on a specific ndoe
func ParseNodeInfoMsg(slurm_response *C.struct_node_info_msg) ([]NodeInfo, error) {
	var node_info []NodeInfo
	cNodeArray := slurm_response.node_array
	length := int(slurm_response.record_count)
	nodes := unsafe.Slice(cNodeArray, length)
	for _, node := range nodes {
		info := NodeInfo{
			Arch:            C.GoString(node.arch),
			TresFmtStr:      C.GoString(node.tres_fmt_str),
			Version:         C.GoString(node.version),
			BcastAddress:    C.GoString(node.bcast_address),
			Boards:          uint16(node.boards),
			BootTime:        time.Unix(int64(node.boot_time), 0),
			ClusterName:     C.GoString(node.cluster_name),
			Cores:           uint16(node.cores),
			CoreSpecCnt:     uint16(node.core_spec_cnt),
			CpuBind:         uint32(node.cpu_bind),
			CpuLoad:         uint32(node.cpu_load),
			FreeMem:         uint64(node.free_mem),
			Cpus:            uint16(node.cpus),
			CpusEfctv:       uint16(node.cpus_efctv),
			CpuSpecList:     C.GoString(node.cpu_spec_list),
			Extra:           C.GoString(node.extra),
			Features:        C.GoString(node.features),
			FeaturesAct:     C.GoString(node.features_act),
			Gres:            C.GoString(node.gres),
			GresDrain:       C.GoString(node.gres_drain),
			GresUsed:        C.GoString(node.gres_used),
			LastBusy:        time.Unix(int64(node.last_busy), 0),
			McsLabel:        C.GoString(node.mcs_label),
			MemSpecLimit:    uint64(node.mem_spec_limit),
			Name:            C.GoString(node.name),
			NextState:       uint32(node.next_state),
			NodeAddr:        C.GoString(node.node_addr),
			NodeHostname:    C.GoString(node.node_hostname),
			NodeState:       uint32(node.node_state),
			OS:              C.GoString(node.os),
			Owner:           uint32(node.owner),
			Partitions:      C.GoString(node.partitions),
			Port:            uint16(node.port),
			RealMemory:      uint64(node.real_memory),
			Comment:         C.GoString(node.comment),
			Reason:          C.GoString(node.reason),
			ReasonTime:      time.Unix(int64(node.reason_time), 0),
			ReasonUid:       uint32(node.reason_uid),
			SlurmdStartTime: time.Unix(int64(node.slurmd_start_time), 0),
			Sockets:         uint16(node.sockets),
			Threads:         uint16(node.threads),
			TmpDisk:         uint32(node.tmp_disk),
			Weight:          uint32(node.weight),
		}
		node_info = append(node_info, info)
	}
	C.slurm_free_node_info_msg(slurm_response)
	return node_info, nil
}

// Ping the SLURM cluster
func RpcPingSlurm() (int, error) {
	exit_reason := C.slurm_ping(C.int(0))
	if exit_reason != 0 {
		return 1, fmt.Errorf("Slurm cluster is unresponsive: exit reason")
	}
	return 0, nil
}
