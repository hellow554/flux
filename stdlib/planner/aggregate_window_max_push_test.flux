package planner_test


import "testing"

input = "
#datatype,string,long,dateTime:RFC3339,string,string,string,double
#group,false,false,false,true,true,true,false
#default,_result,,,,,,
,result,table,_time,_measurement,host,_field,_value
,,0,2018-05-22T19:53:26Z,system,host.local,load1,1.83
,,0,2018-05-22T19:53:36Z,system,host.local,load1,1.72
,,0,2018-05-22T19:53:37Z,system,host.local,load1,1.77
,,0,2018-05-22T19:53:56Z,system,host.local,load1,1.63
,,0,2018-05-22T19:54:06Z,system,host.local,load1,1.91
,,0,2018-05-22T19:54:16Z,system,host.local,load1,1.84

,,1,2018-05-22T19:53:26Z,system,host.local,load3,1.98
,,1,2018-05-22T19:53:36Z,system,host.local,load3,1.97
,,1,2018-05-22T19:53:46Z,system,host.local,load3,1.97
,,1,2018-05-22T19:53:56Z,system,host.local,load3,1.96
,,1,2018-05-22T19:54:06Z,system,host.local,load3,1.99
,,1,2018-05-22T19:54:16Z,system,host.local,load3,1.97

,,2,2018-05-22T19:53:26Z,system,host.local,load5,1.95
,,2,2018-05-22T19:53:36Z,system,host.local,load5,1.92
,,2,2018-05-22T19:53:41Z,system,host.local,load5,1.91
,,2,2018-05-22T19:53:46Z,system,host.local,load5,1.92
,,2,2018-05-22T19:53:56Z,system,host.local,load5,1.89
,,2,2018-05-22T19:54:16Z,system,host.local,load5,1.93
"
output = "
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,string,string,string,double
#group,false,false,true,true,false,true,true,true,false
#default,_result,,,,,,,,
,result,table,_start,_stop,_time,_measurement,host,_field,_value
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:20Z,2018-05-22T19:53:40Z,system,host.local,load1,1.83
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:20Z,2018-05-22T19:54:00Z,system,host.local,load1,1.63
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:20Z,2018-05-22T19:54:20Z,system,host.local,load1,1.91
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:20Z,2018-05-22T19:53:40Z,system,host.local,load3,1.98
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:20Z,2018-05-22T19:54:00Z,system,host.local,load3,1.97
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:20Z,2018-05-22T19:54:20Z,system,host.local,load3,1.99
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:20Z,2018-05-22T19:53:40Z,system,host.local,load5,1.95
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:20Z,2018-05-22T19:54:00Z,system,host.local,load5,1.92
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:20Z,2018-05-22T19:54:20Z,system,host.local,load5,1.93
"
aggregate_window_max_fn = (tables=<-) => tables
    |> range(start: 2018-05-22T19:53:00Z, stop: 2018-05-22T19:54:20Z)
    |> aggregateWindow(every: 20s, fn: max)

test aggregate_window_max_pushdown = () => ({input: testing.loadStorage(csv: input), want: testing.loadMem(csv: output), fn: aggregate_window_max_fn})
