package appconfig

import (
	"encoding/json"
	"testing"
)

func TestCollectionSS(t *testing.T) {
	var kMsg KafkaMSG
	raw := []byte(rawKafkaMsg)
	err := json.Unmarshal(raw, &kMsg)
	if err != nil {
		t.Fatalf("error marshaling raw kafka msg")
	}
	kSha := kMsg.SHA()
	stateFile, err := kMsg.StateFile()
	if err != nil {
		t.Fatalf("error converting kafka message into statefile")
	}
	tps := stateFile.Types()
	ads := stateFile.ADs()
	keys := stateFile.Keys()
	vals := stateFile.Values()
	var easi string
	var node string
	if len(stateFile.Get(`easi`)) > 0 {
		easi = stateFile.Get(`easi`)[0]
	}
	if len(stateFile.Get(`node`)) > 0 {
		node = stateFile.Get(`node`)[0]
	}
	t.Logf("ads  :\t%v\n", len(ads))
	t.Logf("types:\t%v\n", len(tps))
	t.Logf("keys :\t%v\n", len(keys))
	t.Logf("vals :\t%v\n", len(vals))
	t.Logf("kSHA :\t%v\n", kSha)
	t.Logf("easi :\t%v\n", stateFile.Get(`easi`))
	t.Logf("node :\t%v\n", stateFile.Get(`node`))
	switch {
	case len(ads) != 2:
		t.Fatalf("incorrect number of appdomains, expected %v, got %v", 2, len(ads))
	case len(tps) != 3:
		t.Fatalf("incorrect number of types, expected %v, got %v", 3, len(tps))
	case len(keys) != 25:
		t.Fatalf("incorrect number of keys, expected %v, got %v", 25, len(keys))
	case len(vals) != 23:
		t.Fatalf("incorrect number of values, expected %v, got %v", 23, len(vals))
	case easi != `srv-wm-app-packapi`:
		t.Fatalf("incorrect value for easi, expected %v, got %q", `srv-wm-app-packapi`, easi)
	case node != `srv24w0m15`:
		t.Fatalf("incorrect value for easi, expected %v, got %q", `srv24w0m15`, node)
	}
	SS := make(SavedState, 1)
	sf, err := kMsg.SavedFile()
	if err != nil {
		t.Fatalf("error converting kafka message into savedfile")
	}
	SS[0] = sf
	sSha := sf.SHA()
	envs, asis, easis, easins, nodes := SS.GetAll()
	t.Logf("envs   :\t%v\n", len(envs))
	t.Logf("asis   :\t%v\n", len(asis))
	t.Logf("easis  :\t%v\n", len(easis))
	t.Logf("easins :\t%v\n", len(easins))
	t.Logf("nodes  :\t%v\n", len(nodes))
	t.Logf("sSHA :\t%v\n", sSha)
	t.Logf("easi :\t%v\n", sf.EASI)
	t.Logf("node :\t%v\n", sf.Node)
	switch {
	case !sf.HasNode(node):
		t.Fatalf("savedfile did not have expected node value")
	case sSha != kSha:
		t.Fatalf("sha1 values for savedfile do not match kafka message, expected %v received %v", kSha, sSha)
	case len(envs) != 1 || len(asis) != 1 || len(easis) != 1 || len(easins) != 1 || len(nodes) != 1:
		t.Fatalf("incorrect number of top level values found")
	}

}

const rawKafkaMsg = `{"@timestamp":"2019-10-24T21:03:12.009Z","@metadata":{"beat":"filebeat","type":"doc","version":"6.7.2","topic":"srv-appconfig-event-json"},"easi":"srv:wm:app:packapi","host":{"name":"srv24w0m15.example.com"},"log":"appconfig-install.state.json","message":"{\"data\": [{\"appdomain\": null, \"k\": \"node\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"environment\", \"type\": \"simple\", \"v\": \"srv24w0m15\"}, {\"appdomain\": null, \"k\": \"operatingsystemrelease\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"facter\", \"type\": \"simple\", \"v\": \"7.6.1810\"}, {\"appdomain\": null, \"k\": \"packapi\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"etmeta\", \"type\": \"simple\", \"v\": \"sit20191024.103-0\"}, {\"appdomain\": null, \"k\": \"easi\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"environment\", \"type\": \"simple\", \"v\": \"srv-wm-app-packapi\"}, {\"appdomain\": null, \"k\": \"uptime_days\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"facter\", \"type\": \"simple\", \"v\": \"17\"}, {\"appdomain\": null, \"k\": \"appdomain\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"etmeta\", \"type\": \"simple\", \"v\": \"srv1m7\"}, {\"appdomain\": null, \"k\": \"processorcount\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"facter\", \"type\": \"simple\", \"v\": \"2\"}, {\"appdomain\": null, \"k\": \"memorysize_mb\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"facter\", \"type\": \"simple\", \"v\": \"3789.76\"}, {\"appdomain\": null, \"k\": \"timezone\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"facter\", \"type\": \"simple\", \"v\": \"EDT\"}, {\"appdomain\": \"srv1m7\", \"k\": \"advisorxml\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"etmeta\", \"tpls\": [\"config/advisor.conf\"], \"type\": \"endpoint\", \"v\": \"wmax.srv.example.com:9030:http:srv1m7\"}, {\"appdomain\": null, \"k\": \"ports__ADVISOR_HTTP_PORT\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"default\", \"tpls\": [\"config/advisor.conf\"], \"type\": \"parameter\", \"v\": \"8081\"}, {\"appdomain\": null, \"k\": \"endpoint__advisorxml__port\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"etmeta\", \"tpls\": [\"config/advisor.conf\"], \"type\": \"parameter\", \"v\": \"9030\"}, {\"appdomain\": null, \"k\": \"ports__ADVISOR_GRPC_PORT\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"default\", \"tpls\": [\"config/envoy.yaml\"], \"type\": \"parameter\", \"v\": \"9081\"}, {\"appdomain\": null, \"k\": \"ports__AGGREGATOR_HTTP_PORT\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"default\", \"tpls\": [\"config/aggregator.conf\"], \"type\": \"parameter\", \"v\": \"8080\"}, {\"appdomain\": null, \"k\": \"properties__deq-ack-timeout\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"default\", \"tpls\": [\"config/advisor.conf\"], \"type\": \"parameter\", \"v\": \"30\"}, {\"appdomain\": null, \"k\": \"environment__e_ir\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"environment\", \"tpls\": [\"config/envoy.yaml\"], \"type\": \"parameter\", \"v\": \"/example/srv-wm-app-packapi\"}, {\"appdomain\": null, \"k\": \"ports__HEALTHCHECK_PORT\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"default\", \"tpls\": [\"opt/tools/monitor.cfg\"], \"type\": \"parameter\", \"v\": \"8000\"}, {\"appdomain\": null, \"k\": \"ports__BROKER_GRPC_PORT\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"default\", \"tpls\": [\"config/envoy.yaml\"], \"type\": \"parameter\", \"v\": \"9082\"}, {\"appdomain\": null, \"k\": \"ports__BROKER_HTTP_PORT\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"default\", \"tpls\": [\"config/broker.conf\"], \"type\": \"parameter\", \"v\": \"8082\"}, {\"appdomain\": null, \"k\": \"environment__e_node\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"environment\", \"tpls\": [\"opt/tools/monitor.cfg\"], \"type\": \"parameter\", \"v\": \"srv24w0m15\"}, {\"appdomain\": null, \"k\": \"environment__e_envoy_root\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"environment\", \"tpls\": [\"config/supervisor/envoy.conf\"], \"type\": \"parameter\", \"v\": \"/example/srv-wm-app-packapi/packages/envoy\"}, {\"appdomain\": null, \"k\": \"endpoint__advisorxml__endpoint\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"etmeta\", \"tpls\": [\"config/advisor.conf\"], \"type\": \"parameter\", \"v\": \"wmax.srv.example.com\"}, {\"appdomain\": null, \"k\": \"environment__e_packapi_root\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"environment\", \"tpls\": [\"config/supervisor/grpc.conf\"], \"type\": \"parameter\", \"v\": \"/example/srv-wm-app-packapi/packages/packapi\"}, {\"appdomain\": null, \"k\": \"ports__AGGREGATOR_GRPC_PORT\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"default\", \"tpls\": [\"config/envoy.yaml\"], \"type\": \"parameter\", \"v\": \"9080\"}, {\"appdomain\": null, \"k\": \"ports__ENVOY_HTTP_PORT\", \"pkg\": \"packapi-sit20191024.103-0\", \"src\": \"appconfig\", \"tpls\": [\"config/envoy.yaml\"], \"type\": \"parameter\", \"v\": \"8000\"}], \"dttm\": 1571950979.575358}","offset":92453,"node":"srv24w0m15","datacenter":"m15","input":{"type":"log"},"source":"/example/srv-wm-app-packapi/logs/appconfig-install.state.json","prospector":{"type":"log"},"env":"srv","workgroup":"w05","pipeline":{"topic":"srv-appconfig-event-json","source":"filebeat"},"streamSource":"/opt/streams/source/filebeat/appconfigjson.hcl","asi":"wm:app:packapi","beat":{"name":"srv24w0m15.example.com","hostname":"srv24w0m15.example.com","version":"6.7.2"}}`
