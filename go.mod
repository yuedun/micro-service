module micro-service

go 1.13

require (
	github.com/DataDog/datadog-go v3.3.1+incompatible // indirect
	github.com/coredns/coredns v1.6.6 // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/envoyproxy/go-control-plane v0.9.1 // indirect
	github.com/gogo/googleapis v1.3.1 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.4
	github.com/hashicorp/consul v1.6.2 // indirect
	github.com/hashicorp/go-discover v0.0.0-20200106195200-2f5d876fe697 // indirect
	github.com/hashicorp/go-memdb v1.0.4 // indirect
	github.com/hashicorp/hil v0.0.0-20190212132231-97b3a9cdfa93 // indirect
	github.com/hashicorp/raft-boltdb v0.0.0-20191021154308-4207f1bf0617 // indirect
	github.com/hashicorp/vault v1.3.1 // indirect
	github.com/hashicorp/yamux v0.0.0-20190923154419-df201c70410d // indirect
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.2.0
	github.com/micro/protoc-gen-micro/v2 v2.0.0 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/prometheus/client_golang v1.3.0 // indirect
	github.com/russross/blackfriday v0.0.0-20170610170232-067529f716f4 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shirou/gopsutil v2.19.12+incompatible // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.27.1 // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.54.0
	cloud.google.com/go/bigquery => github.com/googleapis/google-cloud-go/bigquery v1.5.0
	cloud.google.com/go/datastore => github.com/googleapis/google-cloud-go/datastore v1.1.0
	cloud.google.com/go/pubsub => github.com/googleapis/google-cloud-go/pubsub v1.3.0
	cloud.google.com/go/storage => github.com/googleapis/google-cloud-go/storage v1.6.0
	github.com/docker/docker => github.com/docker/docker v1.13.1
	go.opencensus.io => github.com/census-instrumentation/opencensus-go v0.19.0
	go.uber.org/multierr => github.com/uber-go/multierr v1.1.0
	go.uber.org/zap => github.com/uber-go/zap v1.9.1
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200302210943-78000ba7a073
	golang.org/x/exp => github.com/golang/exp v0.0.0-20200228211341-fcea875c7e85
	golang.org/x/image => github.com/golang/image v0.0.0-20200119044424-58c23975cae1
	golang.org/x/lint => github.com/golang/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20200222142934-3c8601c510d0
	golang.org/x/mod => github.com/golang/mod v0.2.0
	golang.org/x/net => github.com/golang/net v0.0.0-20200301022130-244492dfa37a
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200302150141-5c8b2ff67527
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20191024005414-555d28b269f0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20200305224536-de023d59a5d1
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191204190536-9bdfabe68543
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.0.0-20181220000619-583d854617af
	google.golang.org/appengine => github.com/golang/appengine v1.6.5
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20200305110556-506484158171
	google.golang.org/grpc => github.com/grpc/grpc-go v1.25.1
	gopkg.in/alecthomas/kingpin.v2 => github.com/alecthomas/kingpin v2.2.6+incompatible
	gopkg.in/mgo.v2 => github.com/go-mgo/mgo v0.0.0-20180705113604-9856a29383ce
	gopkg.in/vmihailenco/msgpack.v2 => github.com/vmihailenco/msgpack v2.9.1+incompatible
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
	labix.org/v2/mgo => github.com/go-mgo/mgo v0.0.0-20160801194620-b6121c6199b7
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
)
