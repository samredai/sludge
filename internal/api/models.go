package api

import (
	"time"
)

type JobDetails struct {
	ID                int       `yaml:"Job ID"`
	Account           string    `yaml:"Account"`
	Username          string    `yaml:"Username"`
	UserID            int       `yaml:"User ID"`
	WorkDir           string    `yaml:"Working Directory"`
	JobName           string    `yaml:"Job Name"`
	SubmitTime        time.Time `yaml:"Submitted Time"`
	StartTime         time.Time `yaml:"Start Time"`
	LastModTime       time.Time `yaml:"Last Modified Time"`
	SystemComment     string    `yaml:"System Comment"`
	TimeLimit         int       `yaml:"Time Limit"`
	TimeMin           int       `yaml:"Minimum Allowed Time"`
	ThreadsPerCore    int       `yaml:"Threads Per Core"`
	StateReason       int       `yaml:"State Reason"`
	StateDescription  string    `yaml:"State Description"`
	StdErr            string    `yaml:"Standard Error File"`
	StdIn             string    `yaml:"Standard In File"`
	StdOut            string    `yaml:"Standard Out File"`
	ShowFlags         int       `yaml:"Show Flags"`
	Shared            int       `yaml:"Shared"`
	ScheduledNodes    string    `yaml:"Scheduled Nodes"`
	ReservationName   string    `yaml:"Reservation Name"`
	RequiredNodes     string    `yaml:"Required Nodes"`
	Priority          int       `yaml:"Priority"`
	PreSusTime        time.Time `yaml:"Pre-Suspended Time"`
	MinTmpDiskPerNode int       `yaml:"Minimum Temp Disk Per Node"`
	MinCPUsPerNode    int       `yaml:"Minimum CPUs Per Node"`
	MinMemPerNode     int       `yaml:"Minimum Memory Per Node"`
	Prefer            string    `yaml:"Prefer"`
	Partition         string    `yaml:"Partition"`
	NumTasks          int       `yaml:"Number of Tasks"`
	NumNodes          int       `yaml:"Number of Nodes"`
	NumCPUs           int       `yaml:"Number of CPUs"`
	NumTasksPerCore   int       `yaml:"Number of Tasks Per Core"`
	NumTasksPerTres   int       `yaml:"Number of Tasks Per Tres"`
	NumTasksPerNode   int       `yaml:"Number of Tasks Per Node"`
	NumTasksPerSocket int       `yaml:"Number of Tasks Per Socket"`
	NumTasksPerBoard  int       `yaml:"Number of Tasks Per Board"`
	RestartCount      int       `yaml:"Restart Count"`
	Nodes             string    `yaml:"Nodes"`
	Network           string    `yaml:"Network"`
	MemPerTres        string    `yaml:"Memory Per Tres"`
	MaxNodes          int       `yaml:"Maximum Nodes"`
	MaxCPUs           int       `yaml:"Maximum CPUs"`
	MailUser          string    `yaml:"Mail User"`
	MailType          int       `yaml:"Mail Type"`
	Licenses          string    `yaml:"Licenses"`
	LastSchedEval     time.Time `yaml:"Last Schedule Evaluation"`
	JobState          int       `yaml:"Job State"`
	HetJobOffset      int       `yaml:"Het Job Offset"`
	HetJobIdSet       string    `yaml:"Het Job ID Set"`
	HetJobID          int       `yaml:"Het Job ID"`
	GroupID           int       `yaml:"Group ID"`
	GresTotal         string    `yaml:"Gres Total"`
	GresDetailStr     string    `yaml:"Gres Detail"`
	GresDetailCnt     int       `yaml:"Gres Detail Count"`
	FedOriginStr      string    `yaml:"Fed Origin"`
	Features          string    `yaml:"Features"`
	ExitCode          int       `yaml:"Exit Code"`
	ExcludedNodes     string    `yaml:"Excluded Nodes"`
	EndTime           time.Time `yaml:"End Time"`
	EligibleTime      time.Time `yaml:"Eligible Time"`
	DerivedExitCode   int       `yaml:"Derived Exit Code"`
	Dependency        string    `yaml:"Dependency"`
	DelayBoot         int       `yaml:"Delay Boot"`
	Deadline          time.Time `yaml:"Deadline"`
	Cronspec          string    `yaml:"Cron Spec"`
	CpusPerTres       string    `yaml:"CPUs Per Tres"`
	CpuFreqGov        int       `yaml:"CPU Freq Gov"`
	CpuFreqMax        int       `yaml:"CPU Freq Max"`
	CpuFreqMin        int       `yaml:"CPU Freq Min"`
	CpusPerTask       int       `yaml:"CPUs Per Task"`
	CoresPerSocket    int       `yaml:"Cores Per Socket"`
	CoreSpec          int       `yaml:"Core Spec"`
	Contiguous        int       `yaml:"Contiguous"`
	Container         string    `yaml:"Container"`
	Comment           string    `yaml:"Comment"`
	Command           string    `yaml:"Command"`
	ClusterFeatures   string    `yaml:"Cluster Features"`
	Cluster           string    `yaml:"Cluster"`
	BurstBufferState  string    `yaml:"Burst Buffer State"`
	BurstBuffer       string    `yaml:"Burst Buffer"`
	BoardsPerNode     int       `yaml:"Boards Per Node"`
	BitFlags          int       `yaml:"Bit Flags"`
	BatchHost         string    `yaml:"Batch Host"`
	BatchFlag         int       `yaml:"Batch Flag"`
	BatchFeatures     string    `yaml:"Batch Features"`
	AssocId           int       `yaml:"Associated ID"`
	ArrayTaskStr      string    `yaml:"Array Task"`
	ArrayMaxTasks     int       `yaml:"Array Max Tasks"`
	ArrayTaskId       int       `yaml:"Array Task ID"`
	ArrayJobId        int       `yaml:"Array Job ID"`
	ArrayBitmap       int64     `yaml:"Array Bitmap"`
	AllocSid          int       `yaml:"Allocated Sid"`
	AllocNode         string    `yaml:"Allocated Node"`
	AdminComment      string    `yaml:"Admin Comment"`
	AccrueTime        time.Time `yaml:"Accrue Time"`
}

type NodeInfo struct {
	Name            string    `yaml:"Name"`                                 /* node name to slurm */
	TresFmtStr      string    `yaml:"TRES Format String"`                   /* str representing configured TRES on node */
	Version         string    `yaml:"Slurm Version"`                        /* Slurm version number */
	Arch            string    `yaml:"Arch"`                                 /* computer architecture */
	BcastAddress    string    `yaml:"Bcast Address"`                        /* BcastAddr (optional) */
	Boards          uint16    `yaml:"Boards"`                               /* total number of boards per node  */
	BootTime        time.Time `yaml:"Boot Time"`                            /* time of node boot */
	ClusterName     string    `yaml:"Cluster Name"`                         /* Cluster name ONLY set in federation */
	Cores           uint16    `yaml:"Cores"`                                /* number of cores per socket       */
	CoreSpecCnt     uint16    `yaml:"Cores Spec Count"`                     /* number of specialized cores on node */
	CpuBind         uint32    `yaml:"Cpu Bind"`                             /* Default task binding */
	CpuLoad         uint32    `yaml:"Cpu Load"`                             /* CPU load * 100 */
	FreeMem         uint64    `yaml:"Free Memory (MB)"`                     /* free memory in MiB */
	Cpus            uint16    `yaml:"CPU Count"`                            /* configured count of cpus running on the node */
	CpusEfctv       uint16    `yaml:"Effective CPU Count"`                  /* count of effective cpus on the node. i.e cpus minus specialized cpus */
	CpuSpecList     string    `yaml:"Specialized CPUs"`                     /* node's specialized cpus */
	Extra           string    `yaml:"Extra"`                                /* Extra information about the node */
	Features        string    `yaml:"Features"`                             /* list of a node's available features */
	FeaturesAct     string    `yaml:"Active Features"`                      /* list of a node's current active features, Same as "features" if NULL */
	Gres            string    `yaml:"Gres"`                                 /* list of a node's generic resources */
	GresDrain       string    `yaml:"Drained Gres"`                         /* list of drained GRES */
	GresUsed        string    `yaml:"Gres Used"`                            /* list of GRES in current use */
	LastBusy        time.Time `yaml:"Last Busy"`                            /* time node was last busy (i.e. no jobs) */
	McsLabel        string    `yaml:"MCS Label"`                            /* mcs label if mcs plugin in use */
	MemSpecLimit    uint64    `yaml:"Memory Limit for Specialization (MB)"` /* MB memory limit for specialization */
	NextState       uint32    `yaml:"Next State"`                           /* state after reboot (enum node_states) */
	NodeAddr        string    `yaml:"Node Address"`                         /* communication name (optional) */
	NodeHostname    string    `yaml:"Node Hostname"`                        /* node's hostname (optional) */
	NodeState       uint32    `yaml:"Node State"`                           /* see enum node_states */
	OS              string    `yaml:"OS"`                                   /* operating system currently running */
	Owner           uint32    `yaml:"Owner"`                                /* User allowed to use this node or NO_VAL */
	Partitions      string    `yaml:"Partitions"`                           /* Comma separated list of partitions containing this node, NOT supplied by slurmctld, but populated by scontrol */
	Port            uint16    `yaml:"Port"`                                 /* TCP port number of the slurmd */
	RealMemory      uint64    `yaml:"Real Memory (MB)"`                     /* configured MB of real memory on the node */
	Comment         string    `yaml:"Comment"`                              /* arbitrary comment */
	Reason          string    `yaml:"Reason"`                               /* reason for node being DOWN or DRAINING */
	ReasonTime      time.Time `yaml:"Reason Time"`                          /* Time stamp when reason was set, ignore if no reason is set. */
	ReasonUid       uint32    `yaml:"Reason UID"`                           /* User that set the reason, ignore if no reason is set. */
	SlurmdStartTime time.Time `yaml:"Slurmd Start Time"`                    /* time of slurmd startup */
	Sockets         uint16    `yaml:"Sockets"`                              /* total number of sockets per node */
	Threads         uint16    `yaml:"Threads"`                              /* number of threads per core */
	TmpDisk         uint32    `yaml:"Tmp Disk (MB)"`                        /* configured MB of total disk in TMP_FS */
	Weight          uint32    `yaml:"Weight"`                               /* arbitrary priority of node for scheduling */
}

type ClusterStats struct {
	PartsPacked          int       `yaml:"Parts Packed"`
	ReqTime              time.Time `yaml:"Req Time"`
	ReqTimeStart         time.Time `yaml:"Req Time Start"`
	ServerThreadCount    int       `yaml:"Server Thread Count"`
	AgentQueueSize       int       `yaml:"Agent Queue Size"`
	AgentCount           int       `yaml:"Agent Count"`
	AgentThreadCount     int       `yaml:"Agent Thread Count"`
	DbdAgentQueueSize    int       `yaml:"Dbd Agent Queue Size"`
	GetTimeOfDayLatency  int       `yaml:"Get Time Of Day Latency"`
	ScheduleCycleMax     int       `yaml:"Schedule Cycle Max"`
	ScheduleCycleLast    int       `yaml:"Schedule Cycle Last"`
	ScheduleCycleSum     int       `yaml:"Schedule Cycle Sum"`
	ScheduleCycleCounter int       `yaml:"Schedule Cycle Counter"`
	ScheduleCycleDepth   int       `yaml:"Schedule Cycle Depth"`
	ScheduleQueueLen     int       `yaml:"Schedule Queue Len"`
	JobsSubmitted        int       `yaml:"Jobs Submitted"`
	JobsStarted          int       `yaml:"Jobs Started"`
	JobsCompleted        int       `yaml:"Jobs Completed"`
	JobsCanceled         int       `yaml:"Jobs Canceled"`
	JobsFailed           int       `yaml:"Jobs Failed"`
	JobsPending          int       `yaml:"Jobs Pending"`
	JobsRunning          int       `yaml:"Jobs Running"`
	JobStatesTs          time.Time `yaml:"Job States Ts"`
	BfBackfilledJobs     int       `yaml:"Bf Backfilled Jobs"`
	BfLastBackfilledJobs int       `yaml:"Bf Last Backfilled Jobs"`
	BfBackfilledHetJobs  int       `yaml:"Bf Backfilled Het Jobs"`
	BfCycleCounter       int       `yaml:"Bf Cycle Counter"`
	BfCycleSum           int64     `yaml:"Bf Cycle Sum"`
	BfCycleLast          int       `yaml:"Bf Cycle Last"`
	BfCycleMax           int       `yaml:"Bf Cycle Max"`
	BfLastDepth          int       `yaml:"Bf Last Depth"`
	BfLastDepthTry       int       `yaml:"Bf Last Depth Try"`
	BfDepthSum           int       `yaml:"Bf Depth Sum"`
	BfDepthTrySum        int       `yaml:"Bf Depth Try Sum"`
	BfQueueLen           int       `yaml:"Bf Queue Len"`
	BfQueueLenSum        int       `yaml:"Bf Queue Len Sum"`
	BfTableSize          int       `yaml:"Bf Table Size"`
	BfTableSizeSum       int       `yaml:"Bf Table Size Sum"`
	BfWhenLastCycle      time.Time `yaml:"Bf When Last Cycle"`
	BfActive             int       `yaml:"Bf Active"`
	RpcTypeSize          int       `yaml:"Rpc Type Size"`
	RpcTypeID            uint16    `yaml:"Rpc Type ID"`
	RpcTypeCnt           int       `yaml:"Rpc Type Cnt"`
	RpcTypeTime          int64     `yaml:"Rpc Type Time"`
	RpcUserSize          int       `yaml:"Rpc User Size"`
	RpcUserID            uint      `yaml:"Rpc User ID"`
	RpcUserCnt           uint      `yaml:"Rpc User Cnt"`
	RpcUserTime          int64     `yaml:"Rpc User Time"`
	RpcQueueTypeCount    int       `yaml:"Rpc Queue Type Count"`
	RpcQueueTypeID       int       `yaml:"Rpc Queue Type ID"`
	RpcQueueCount        int       `yaml:"Rpc Queue Count"`
	RpcDumpCount         int       `yaml:"Rpc Dump Count"`
	RpcDumpTypes         int       `yaml:"Rpc Dump Types"`
	RpcDumpHostList      string    `yaml:"Rpc Dump Host List"`
}

const (
	Healthy int = iota
	Down
)

type SlurmPingResult struct {
	Status        int
	PingTime      time.Time
	PingStartTime time.Time
	PingDuration  time.Duration
}
