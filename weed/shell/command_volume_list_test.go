package shell

import (
	"github.com/chrislusf/seaweedfs/weed/storage/types"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"

	"github.com/chrislusf/seaweedfs/weed/pb/master_pb"
)

func TestParsing(t *testing.T) {
	topo := parseOutput(topoData)

	assert.Equal(t, 5, len(topo.DataCenterInfos))

}

func parseOutput(output string) *master_pb.TopologyInfo {
	lines := strings.Split(output, "\n")
	var topo *master_pb.TopologyInfo
	var dc *master_pb.DataCenterInfo
	var rack *master_pb.RackInfo
	var dn *master_pb.DataNodeInfo
	var disk *master_pb.DiskInfo
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "Topology":
			if topo == nil {
				topo = &master_pb.TopologyInfo{}
			}
		case "DataCenter":
			if dc == nil {
				dc = &master_pb.DataCenterInfo{
					Id: parts[1],
				}
				topo.DataCenterInfos = append(topo.DataCenterInfos, dc)
			} else {
				dc = nil
			}
		case "Rack":
			if rack == nil {
				rack = &master_pb.RackInfo{
					Id: parts[1],
				}
				dc.RackInfos = append(dc.RackInfos, rack)
			} else {
				rack = nil
			}
		case "DataNode":
			if dn == nil {
				dn = &master_pb.DataNodeInfo{
					Id:        parts[1],
					DiskInfos: make(map[string]*master_pb.DiskInfo),
				}
				rack.DataNodeInfos = append(rack.DataNodeInfos, dn)
			} else {
				dn = nil
			}
		case "Disk":
			if disk == nil {
				diskType := parts[1][:strings.Index(parts[1], "(")]
				maxVolumeCountStr := parts[1][strings.Index(parts[1], "/")+1:]
				maxVolumeCount, _ := strconv.Atoi(maxVolumeCountStr)
				disk = &master_pb.DiskInfo{
					Type:           diskType,
					MaxVolumeCount: int64(maxVolumeCount),
				}
				dn.DiskInfos[types.ToDiskType(diskType).String()] = disk
			} else {
				disk = nil
			}
		case "volume":
			volumeLine := line[len("volume "):]
			volume := &master_pb.VolumeInformationMessage{}
			proto.UnmarshalText(volumeLine, volume)
			disk.VolumeInfos = append(disk.VolumeInfos, volume)
		}
	}

	return topo
}

const topoData = `
Topology volumeSizeLimit:1024 MB hdd(volume:760/7280 active:760 free:6520 remote:0)
  DataCenter dc1 hdd(volume:0/0 active:0 free:0 remote:0)
    Rack DefaultRack hdd(volume:0/0 active:0 free:0 remote:0)
    Rack DefaultRack total size:0 file_count:0 
  DataCenter dc1 total size:0 file_count:0 
  DataCenter dc2 hdd(volume:86/430 active:86 free:344 remote:0)
    Rack rack1 hdd(volume:50/240 active:50 free:190 remote:0)
      DataNode 192.168.1.4:8080 hdd(volume:50/240 active:50 free:190 remote:0)
        Disk hdd(volume:50/240 active:50 free:190 remote:0)
          volume id:15  size:1115965064  collection:"collection0"  file_count:83  replica_placement:100  version:3  modified_at_second:1609923671 
          volume id:21  size:1097631536  collection:"collection0"  file_count:82  delete_count:7  deleted_byte_count:68975485  replica_placement:100  version:3  modified_at_second:1609929578 
          volume id:22  size:1086828272  collection:"collection0"  file_count:75  replica_placement:100  version:3  modified_at_second:1609930001 
          volume id:23  size:1076380216  collection:"collection0"  file_count:68  replica_placement:100  version:3  modified_at_second:1609930434 
          volume id:24  size:1074139776  collection:"collection0"  file_count:90  replica_placement:100  version:3  modified_at_second:1609930909 
          volume id:25  size:690757512  collection:"collection0"  file_count:38  replica_placement:100  version:3  modified_at_second:1611144216 
          volume id:27  size:298886792  file_count:1608  replica_placement:100  version:3  modified_at_second:1615632482 
          volume id:28  size:308919192  file_count:1591  delete_count:1  deleted_byte_count:125280  replica_placement:100  version:3  modified_at_second:1615631762 
          volume id:29  size:281582680  file_count:1537  replica_placement:100  version:3  modified_at_second:1615629422 
          volume id:30  size:289466144  file_count:1566  delete_count:1  deleted_byte_count:124972  replica_placement:100  version:3  modified_at_second:1615632422 
          volume id:31  size:273363256  file_count:1498  replica_placement:100  version:3  modified_at_second:1615631642 
          volume id:33  size:1130226400  collection:"collection1"  file_count:7322  delete_count:172  deleted_byte_count:45199399  replica_placement:100  version:3  modified_at_second:1615618789 
          volume id:38  size:1075545744  collection:"collection1"  file_count:13324  delete_count:100  deleted_byte_count:25223906  replica_placement:100  version:3  modified_at_second:1615569830 
          volume id:51  size:1076796120  collection:"collection1"  file_count:10550  delete_count:39  deleted_byte_count:12723654  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615547786 
          volume id:52  size:1083529728  collection:"collection1"  file_count:10128  delete_count:32  deleted_byte_count:10608391  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615599195 
          volume id:54  size:1045022344  collection:"collection1"  file_count:9408  delete_count:30  deleted_byte_count:15132106  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630812 
          volume id:63  size:956941112  collection:"collection1"  file_count:8271  delete_count:32  deleted_byte_count:15876189  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632036 
          volume id:69  size:869213648  collection:"collection1"  file_count:7293  delete_count:102  deleted_byte_count:30643207  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630534 
          volume id:74  size:957046128  collection:"collection1"  file_count:6982  delete_count:258  deleted_byte_count:73054259  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615631460 
          volume id:80  size:827912928  collection:"collection1"  file_count:6914  delete_count:17  deleted_byte_count:5689635  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615631157 
          volume id:84  size:873121856  collection:"collection1"  file_count:8200  delete_count:13  deleted_byte_count:3131676  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631161 
          volume id:85  size:1023869320  collection:"collection1"  file_count:7788  delete_count:234  deleted_byte_count:78037967  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631723 
          volume id:97  size:1053112992  collection:"collection1"  file_count:6789  delete_count:50  deleted_byte_count:38894001  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631193 
          volume id:98  size:1077836440  collection:"collection1"  file_count:7605  delete_count:202  deleted_byte_count:73180379  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615523691 
          volume id:105  size:1073996824  collection:"collection1"  file_count:6872  delete_count:20  deleted_byte_count:14482293  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615499757 
          volume id:106  size:1075458664  collection:"collection1"  file_count:7182  delete_count:307  deleted_byte_count:69349053  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615598137 
          volume id:112  size:1076392512  collection:"collection1"  file_count:8291  delete_count:156  deleted_byte_count:74120183  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615569823 
          volume id:116  size:1074489504  collection:"collection1"  file_count:9981  delete_count:174  deleted_byte_count:53998777  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615611565 
          volume id:119  size:1075940104  collection:"collection1"  file_count:9003  delete_count:12  deleted_byte_count:9128155  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615573878 
          volume id:128  size:1074874632  collection:"collection1"  file_count:9821  delete_count:148  deleted_byte_count:43633334  replica_placement:100  version:3  modified_at_second:1615602670 
          volume id:133  size:1075952760  collection:"collection1"  file_count:9538  delete_count:74  deleted_byte_count:19558008  replica_placement:100  version:3  modified_at_second:1615584779 
          volume id:136  size:1074433552  collection:"collection1"  file_count:9593  delete_count:72  deleted_byte_count:26912512  replica_placement:100  version:3  modified_at_second:1615376036 
          volume id:138  size:1074465744  collection:"collection1"  file_count:10120  delete_count:55  deleted_byte_count:15875438  replica_placement:100  version:3  modified_at_second:1615572231 
          volume id:140  size:1076203744  collection:"collection1"  file_count:11219  delete_count:57  deleted_byte_count:19864498  replica_placement:100  version:3  modified_at_second:1615571947 
          volume id:144  size:1074549720  collection:"collection1"  file_count:8780  delete_count:50  deleted_byte_count:52475146  replica_placement:100  version:3  modified_at_second:1615573451 
          volume id:161  size:1077397192  collection:"collection1"  file_count:9988  delete_count:28  deleted_byte_count:12509164  replica_placement:100  version:3  modified_at_second:1615631452 
          volume id:173  size:1074154704  collection:"collection1"  file_count:30884  delete_count:34  deleted_byte_count:2578509  replica_placement:100  version:3  modified_at_second:1615591904 
          volume id:174  size:1073824232  collection:"collection1"  file_count:30689  delete_count:36  deleted_byte_count:2160116  replica_placement:100  version:3  modified_at_second:1615598914 
          volume id:197  size:1075423240  collection:"collection1"  file_count:16473  delete_count:15  deleted_byte_count:12552442  replica_placement:100  version:3  modified_at_second:1615485254 
          volume id:219  size:1092298904  collection:"collection1"  file_count:3193  delete_count:17  deleted_byte_count:2047576  replica_placement:100  version:3  modified_at_second:1615579316 
          volume id:263  size:1077167352  collection:"collection2"  file_count:20227  delete_count:4  deleted_byte_count:97887  replica_placement:100  version:3  modified_at_second:1614871567 
          volume id:272  size:1076146040  collection:"collection2"  file_count:21034  delete_count:2  deleted_byte_count:216564  replica_placement:100  version:3  modified_at_second:1614884139 
          volume id:291  size:1076256760  collection:"collection2"  file_count:28301  delete_count:5  deleted_byte_count:116027  replica_placement:100  version:3  modified_at_second:1614904924 
          volume id:299  size:1075147824  collection:"collection2"  file_count:22927  delete_count:4  deleted_byte_count:345569  replica_placement:100  version:3  modified_at_second:1614918454 
          volume id:301  size:1074655600  collection:"collection2"  file_count:22543  delete_count:6  deleted_byte_count:136968  replica_placement:100  version:3  modified_at_second:1614918378 
          volume id:302  size:1077559792  collection:"collection2"  file_count:23124  delete_count:7  deleted_byte_count:293111  replica_placement:100  version:3  modified_at_second:1614925500 
          volume id:339  size:1078402392  collection:"collection2"  file_count:22309  replica_placement:100  version:3  modified_at_second:1614969996 
          volume id:345  size:1074560760  collection:"collection2"  file_count:22117  delete_count:2  deleted_byte_count:373286  replica_placement:100  version:3  modified_at_second:1614977458 
          volume id:355  size:1075239792  collection:"collection2"  file_count:22244  delete_count:1  deleted_byte_count:23282  replica_placement:100  version:3  modified_at_second:1614992157 
          volume id:373  size:1080928000  collection:"collection2"  file_count:22617  delete_count:4  deleted_byte_count:91849  replica_placement:100  version:3  modified_at_second:1615016877 
        Disk hdd total size:48630015544 file_count:537880 deleted_file:2580 deleted_bytes:929560253 
      DataNode 192.168.1.4:8080 total size:48630015544 file_count:537880 deleted_file:2580 deleted_bytes:929560253 
    Rack rack1 total size:48630015544 file_count:537880 deleted_file:2580 deleted_bytes:929560253 
    Rack rack2 hdd(volume:36/190 active:36 free:154 remote:0)
      DataNode 192.168.1.2:8080 hdd(volume:36/190 active:36 free:154 remote:0)
        Disk hdd(volume:36/190 active:36 free:154 remote:0)
          volume id:2  size:289228560  file_count:1640  delete_count:4  deleted_byte_count:480564  replica_placement:100  version:3  compact_revision:6  modified_at_second:1615630622 
          volume id:3  size:308743136  file_count:1638  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632242 
          volume id:4  size:285986968  file_count:1641  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632302 
          volume id:6  size:302411024  file_count:1604  delete_count:2  deleted_byte_count:274587  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631402 
          volume id:7  size:1924728  collection:"collection4"  file_count:15  replica_placement:100  version:3  modified_at_second:1609331040 
          volume id:9  size:77337416  collection:"collection3"  file_count:58  replica_placement:100  version:3  ttl:772  modified_at_second:1615513762 
          volume id:10  size:1212784656  collection:"collection0"  file_count:58  replica_placement:100  version:3  modified_at_second:1609814550 
          volume id:12  size:1110923848  collection:"collection0"  file_count:45  replica_placement:100  version:3  modified_at_second:1609819732 
          volume id:13  size:1184910656  collection:"collection0"  file_count:47  replica_placement:100  version:3  modified_at_second:1609827837 
          volume id:14  size:1107475720  collection:"collection0"  file_count:80  delete_count:3  deleted_byte_count:6870  replica_placement:100  version:3  modified_at_second:1612956980 
          volume id:16  size:1113666104  collection:"collection0"  file_count:73  delete_count:5  deleted_byte_count:6318  replica_placement:100  version:3  modified_at_second:1612957007 
          volume id:17  size:1095115800  collection:"collection0"  file_count:83  delete_count:3  deleted_byte_count:7099  replica_placement:100  version:3  modified_at_second:1612957000 
          volume id:21  size:1097631664  collection:"collection0"  file_count:82  delete_count:11  deleted_byte_count:68985100  replica_placement:100  version:3  modified_at_second:1612957007 
          volume id:56  size:1001897616  collection:"collection1"  file_count:8762  delete_count:37  deleted_byte_count:65375405  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615632014 
          volume id:81  size:880693104  collection:"collection1"  file_count:7481  delete_count:236  deleted_byte_count:80386421  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615631396 
          volume id:104  size:1076383624  collection:"collection1"  file_count:7663  delete_count:184  deleted_byte_count:100728071  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615602658 
          volume id:107  size:1073811840  collection:"collection1"  file_count:7436  delete_count:168  deleted_byte_count:57747484  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615293569 
          volume id:113  size:1076709184  collection:"collection1"  file_count:9355  delete_count:177  deleted_byte_count:59796765  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615569822 
          volume id:139  size:1074163936  collection:"collection1"  file_count:9315  delete_count:42  deleted_byte_count:10630966  replica_placement:100  version:3  modified_at_second:1615571946 
          volume id:151  size:1098659752  collection:"collection1"  file_count:10808  delete_count:24  deleted_byte_count:7088102  replica_placement:100  version:3  modified_at_second:1615586389 
          volume id:155  size:1075140688  collection:"collection1"  file_count:10882  delete_count:32  deleted_byte_count:9076141  replica_placement:100  version:3  modified_at_second:1615606614 
          volume id:167  size:1073958176  collection:"collection1"  file_count:25229  delete_count:48  deleted_byte_count:25871565  replica_placement:100  version:3  modified_at_second:1615602669 
          volume id:177  size:1074120216  collection:"collection1"  file_count:22293  delete_count:16  deleted_byte_count:3803952  replica_placement:100  version:3  modified_at_second:1615516892 
          volume id:179  size:1074313920  collection:"collection1"  file_count:21829  delete_count:24  deleted_byte_count:45552859  replica_placement:100  version:3  modified_at_second:1615580308 
          volume id:182  size:1076131280  collection:"collection1"  file_count:31987  delete_count:21  deleted_byte_count:1452346  replica_placement:100  version:3  modified_at_second:1615568922 
          volume id:215  size:1068268216  collection:"collection1"  file_count:2813  delete_count:10  deleted_byte_count:5676795  replica_placement:100  version:3  modified_at_second:1615586386 
          volume id:217  size:1075381872  collection:"collection1"  file_count:3331  delete_count:14  deleted_byte_count:2009141  replica_placement:100  version:3  modified_at_second:1615401638 
          volume id:283  size:1080178944  collection:"collection2"  file_count:19462  delete_count:7  deleted_byte_count:660407  replica_placement:100  version:3  modified_at_second:1614896626 
          volume id:303  size:1075944504  collection:"collection2"  file_count:22541  delete_count:2  deleted_byte_count:13617  replica_placement:100  version:3  modified_at_second:1614925431 
          volume id:309  size:1075178624  collection:"collection2"  file_count:22692  delete_count:3  deleted_byte_count:171124  replica_placement:100  version:3  modified_at_second:1614931409 
          volume id:323  size:1074608200  collection:"collection2"  file_count:21605  delete_count:4  deleted_byte_count:172090  replica_placement:100  version:3  modified_at_second:1614950526 
          volume id:344  size:1075035448  collection:"collection2"  file_count:21765  delete_count:1  deleted_byte_count:24623  replica_placement:100  version:3  modified_at_second:1614977465 
          volume id:347  size:1075145496  collection:"collection2"  file_count:22178  delete_count:1  deleted_byte_count:79392  replica_placement:100  version:3  modified_at_second:1614984727 
          volume id:357  size:1074276208  collection:"collection2"  file_count:23137  delete_count:4  deleted_byte_count:188487  replica_placement:100  version:3  modified_at_second:1614998792 
          volume id:380  size:1010760456  collection:"collection2"  file_count:14921  delete_count:6  deleted_byte_count:65678  replica_placement:100  version:3  modified_at_second:1615632322 
          volume id:381  size:939292792  collection:"collection2"  file_count:14619  delete_count:2  deleted_byte_count:5119  replica_placement:100  version:3  modified_at_second:1615632324 
        Disk hdd total size:33468194376 file_count:369168 deleted_file:1091 deleted_bytes:546337088 
      DataNode 192.168.1.2:8080 total size:33468194376 file_count:369168 deleted_file:1091 deleted_bytes:546337088 
    Rack rack2 total size:33468194376 file_count:369168 deleted_file:1091 deleted_bytes:546337088 
  DataCenter dc2 total size:82098209920 file_count:907048 deleted_file:3671 deleted_bytes:1475897341 
  DataCenter dc3 hdd(volume:108/1850 active:108 free:1742 remote:0)
    Rack rack3 hdd(volume:108/1850 active:108 free:1742 remote:0)
      DataNode 192.168.1.6:8080 hdd(volume:108/1850 active:108 free:1742 remote:0)
        Disk hdd(volume:108/1850 active:108 free:1742 remote:0)
          volume id:1  size:284685936  file_count:1557  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632062 
          volume id:32  size:281390512  file_count:1496  delete_count:6  deleted_byte_count:546403  replica_placement:100  version:3  modified_at_second:1615632362 
          volume id:47  size:444599784  collection:"collection1"  file_count:709  delete_count:19  deleted_byte_count:11913451  replica_placement:100  version:3  modified_at_second:1615632397 
          volume id:49  size:1078775288  collection:"collection1"  file_count:9636  delete_count:22  deleted_byte_count:5625976  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630446 
          volume id:68  size:898630584  collection:"collection1"  file_count:6934  delete_count:95  deleted_byte_count:27460707  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632284 
          volume id:88  size:1073767976  collection:"collection1"  file_count:14995  delete_count:206  deleted_byte_count:81222360  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615629897 
          volume id:202  size:1077533160  collection:"collection1"  file_count:2847  delete_count:67  deleted_byte_count:65172985  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615588497 
          volume id:203  size:1027316272  collection:"collection1"  file_count:3040  delete_count:11  deleted_byte_count:3993230  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615631728 
          volume id:205  size:1078485304  collection:"collection1"  file_count:2869  delete_count:43  deleted_byte_count:18290259  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615579314 
          volume id:206  size:1082045848  collection:"collection1"  file_count:2979  delete_count:225  deleted_byte_count:88220074  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615564274 
          volume id:209  size:1074083592  collection:"collection1"  file_count:3238  delete_count:4  deleted_byte_count:1494244  replica_placement:100  version:3  modified_at_second:1615419954 
          volume id:211  size:1080610712  collection:"collection1"  file_count:3247  delete_count:7  deleted_byte_count:1891456  replica_placement:100  version:3  modified_at_second:1615269124 
          volume id:212  size:1078293360  collection:"collection1"  file_count:3106  delete_count:6  deleted_byte_count:2085755  replica_placement:100  version:3  modified_at_second:1615586387 
          volume id:213  size:1093587976  collection:"collection1"  file_count:3681  delete_count:12  deleted_byte_count:3138791  replica_placement:100  version:3  modified_at_second:1615586387 
          volume id:214  size:1074486992  collection:"collection1"  file_count:3217  delete_count:10  deleted_byte_count:6392871  replica_placement:100  version:3  modified_at_second:1615586383 
          volume id:216  size:1080073496  collection:"collection1"  file_count:3316  delete_count:4  deleted_byte_count:179819  replica_placement:100  version:3  modified_at_second:1615586387 
          volume id:222  size:1106623104  collection:"collection1"  file_count:3273  delete_count:11  deleted_byte_count:2114627  replica_placement:100  version:3  modified_at_second:1615586243 
          volume id:223  size:1075233064  collection:"collection1"  file_count:2966  delete_count:9  deleted_byte_count:744001  replica_placement:100  version:3  modified_at_second:1615586244 
          volume id:227  size:1106699896  collection:"collection1"  file_count:2827  delete_count:20  deleted_byte_count:5496790  replica_placement:100  version:3  modified_at_second:1615609989 
          volume id:229  size:1109855312  collection:"collection1"  file_count:2857  delete_count:22  deleted_byte_count:2839883  replica_placement:100  version:3  modified_at_second:1615609988 
          volume id:230  size:1080722984  collection:"collection1"  file_count:2898  delete_count:15  deleted_byte_count:3929261  replica_placement:100  version:3  modified_at_second:1615610537 
          volume id:231  size:1112917696  collection:"collection1"  file_count:3151  delete_count:20  deleted_byte_count:2989828  replica_placement:100  version:3  modified_at_second:1615611350 
          volume id:233  size:1080526464  collection:"collection1"  file_count:3136  delete_count:61  deleted_byte_count:17991717  replica_placement:100  version:3  modified_at_second:1615611352 
          volume id:234  size:1073835280  collection:"collection1"  file_count:2965  delete_count:41  deleted_byte_count:4960354  replica_placement:100  version:3  modified_at_second:1615611351 
          volume id:235  size:1075586104  collection:"collection1"  file_count:2767  delete_count:33  deleted_byte_count:3216540  replica_placement:100  version:3  modified_at_second:1615611354 
          volume id:237  size:375722792  collection:"collection1"  file_count:736  delete_count:16  deleted_byte_count:4464870  replica_placement:100  version:3  modified_at_second:1615631727 
          volume id:239  size:426569024  collection:"collection1"  file_count:693  delete_count:19  deleted_byte_count:13020783  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630838 
          volume id:241  size:380217424  collection:"collection1"  file_count:633  delete_count:6  deleted_byte_count:1715768  replica_placement:100  version:3  modified_at_second:1615632006 
          volume id:244  size:1080295352  collection:"collection2"  file_count:10812  delete_count:1  deleted_byte_count:795  replica_placement:100  version:3  modified_at_second:1614852171 
          volume id:245  size:1074597056  collection:"collection2"  file_count:10371  delete_count:3  deleted_byte_count:209701  replica_placement:100  version:3  modified_at_second:1614852093 
          volume id:246  size:1075998648  collection:"collection2"  file_count:10365  delete_count:1  deleted_byte_count:13112  replica_placement:100  version:3  modified_at_second:1614852105 
          volume id:248  size:1084301184  collection:"collection2"  file_count:11217  delete_count:4  deleted_byte_count:746488  replica_placement:100  version:3  modified_at_second:1614856285 
          volume id:249  size:1074819136  collection:"collection2"  file_count:10763  delete_count:2  deleted_byte_count:271699  replica_placement:100  version:3  modified_at_second:1614856230 
          volume id:251  size:1075684488  collection:"collection2"  file_count:10847  replica_placement:100  version:3  modified_at_second:1614856270 
          volume id:252  size:1075065208  collection:"collection2"  file_count:14622  delete_count:2  deleted_byte_count:5228  replica_placement:100  version:3  modified_at_second:1614861196 
          volume id:253  size:1087328816  collection:"collection2"  file_count:14920  delete_count:3  deleted_byte_count:522994  replica_placement:100  version:3  modified_at_second:1614861255 
          volume id:255  size:1079581640  collection:"collection2"  file_count:14877  delete_count:3  deleted_byte_count:101223  replica_placement:100  version:3  modified_at_second:1614861233 
          volume id:256  size:1074283592  collection:"collection2"  file_count:14157  delete_count:1  deleted_byte_count:18156  replica_placement:100  version:3  modified_at_second:1614861100 
          volume id:258  size:1075527216  collection:"collection2"  file_count:18421  delete_count:4  deleted_byte_count:267833  replica_placement:100  version:3  modified_at_second:1614866420 
          volume id:259  size:1075507776  collection:"collection2"  file_count:18079  delete_count:2  deleted_byte_count:71992  replica_placement:100  version:3  modified_at_second:1614866381 
          volume id:264  size:1081624192  collection:"collection2"  file_count:21151  replica_placement:100  version:3  modified_at_second:1614871629 
          volume id:265  size:1076401104  collection:"collection2"  file_count:19932  delete_count:2  deleted_byte_count:160823  replica_placement:100  version:3  modified_at_second:1615629130 
          volume id:266  size:1075617464  collection:"collection2"  file_count:20075  delete_count:1  deleted_byte_count:1039  replica_placement:100  version:3  modified_at_second:1614871526 
          volume id:267  size:1075699544  collection:"collection2"  file_count:21039  delete_count:3  deleted_byte_count:59956  replica_placement:100  version:3  modified_at_second:1614877294 
          volume id:268  size:1074490592  collection:"collection2"  file_count:21698  delete_count:1  deleted_byte_count:33968  replica_placement:100  version:3  modified_at_second:1614877434 
          volume id:269  size:1077552872  collection:"collection2"  file_count:21875  delete_count:4  deleted_byte_count:347272  replica_placement:100  version:3  modified_at_second:1614877481 
          volume id:270  size:1076876568  collection:"collection2"  file_count:22057  delete_count:1  deleted_byte_count:43916  replica_placement:100  version:3  modified_at_second:1614877469 
          volume id:275  size:1078349024  collection:"collection2"  file_count:20808  delete_count:1  deleted_byte_count:1118  replica_placement:100  version:3  modified_at_second:1614884147 
          volume id:277  size:1074956288  collection:"collection2"  file_count:19260  delete_count:2  deleted_byte_count:172356  replica_placement:100  version:3  modified_at_second:1614889988 
          volume id:278  size:1078798640  collection:"collection2"  file_count:20597  delete_count:5  deleted_byte_count:400060  replica_placement:100  version:3  modified_at_second:1614890292 
          volume id:279  size:1077325040  collection:"collection2"  file_count:19671  delete_count:6  deleted_byte_count:379116  replica_placement:100  version:3  modified_at_second:1614890229 
          volume id:280  size:1077432216  collection:"collection2"  file_count:20286  delete_count:1  deleted_byte_count:879  replica_placement:100  version:3  modified_at_second:1614890262 
          volume id:281  size:1077581096  collection:"collection2"  file_count:20206  delete_count:3  deleted_byte_count:143964  replica_placement:100  version:3  modified_at_second:1614890237 
          volume id:284  size:1074533384  collection:"collection2"  file_count:22196  delete_count:4  deleted_byte_count:154683  replica_placement:100  version:3  modified_at_second:1614897231 
          volume id:285  size:1082128688  collection:"collection2"  file_count:21804  delete_count:1  deleted_byte_count:1064  replica_placement:100  version:3  modified_at_second:1614897165 
          volume id:289  size:1075284256  collection:"collection2"  file_count:29342  delete_count:5  deleted_byte_count:100454  replica_placement:100  version:3  modified_at_second:1614904977 
          volume id:290  size:1074723792  collection:"collection2"  file_count:28340  delete_count:4  deleted_byte_count:199064  replica_placement:100  version:3  modified_at_second:1614904924 
          volume id:291  size:1076256768  collection:"collection2"  file_count:28301  delete_count:5  deleted_byte_count:116027  replica_placement:100  version:3  modified_at_second:1614904924 
          volume id:293  size:1075409792  collection:"collection2"  file_count:26063  delete_count:4  deleted_byte_count:183834  replica_placement:100  version:3  modified_at_second:1614912235 
          volume id:294  size:1075444048  collection:"collection2"  file_count:26076  delete_count:4  deleted_byte_count:194914  replica_placement:100  version:3  modified_at_second:1614912220 
          volume id:296  size:1077824032  collection:"collection2"  file_count:26741  delete_count:4  deleted_byte_count:199906  replica_placement:100  version:3  modified_at_second:1614912301 
          volume id:297  size:1080229136  collection:"collection2"  file_count:23409  delete_count:5  deleted_byte_count:46268  replica_placement:100  version:3  modified_at_second:1614918481 
          volume id:298  size:1075410136  collection:"collection2"  file_count:23222  delete_count:2  deleted_byte_count:46110  replica_placement:100  version:3  modified_at_second:1614918474 
          volume id:299  size:1075147936  collection:"collection2"  file_count:22927  delete_count:4  deleted_byte_count:345569  replica_placement:100  version:3  modified_at_second:1614918455 
          volume id:300  size:1076212392  collection:"collection2"  file_count:22892  delete_count:2  deleted_byte_count:61320  replica_placement:100  version:3  modified_at_second:1614918464 
          volume id:301  size:1074655600  collection:"collection2"  file_count:22543  delete_count:6  deleted_byte_count:136968  replica_placement:100  version:3  modified_at_second:1614918378 
          volume id:303  size:1075944480  collection:"collection2"  file_count:22541  delete_count:2  deleted_byte_count:13617  replica_placement:100  version:3  modified_at_second:1614925431 
          volume id:306  size:1074764016  collection:"collection2"  file_count:22939  replica_placement:100  version:3  modified_at_second:1614925462 
          volume id:307  size:1076568000  collection:"collection2"  file_count:23377  delete_count:2  deleted_byte_count:25453  replica_placement:100  version:3  modified_at_second:1614931448 
          volume id:308  size:1074022392  collection:"collection2"  file_count:23086  delete_count:2  deleted_byte_count:2127  replica_placement:100  version:3  modified_at_second:1614931401 
          volume id:309  size:1075178664  collection:"collection2"  file_count:22692  delete_count:3  deleted_byte_count:171124  replica_placement:100  version:3  modified_at_second:1614931409 
          volume id:310  size:1074761528  collection:"collection2"  file_count:21441  delete_count:3  deleted_byte_count:13934  replica_placement:100  version:3  modified_at_second:1614931077 
          volume id:314  size:1074670840  collection:"collection2"  file_count:20964  delete_count:4  deleted_byte_count:304291  replica_placement:100  version:3  modified_at_second:1614937441 
          volume id:315  size:1084153544  collection:"collection2"  file_count:23638  delete_count:2  deleted_byte_count:53956  replica_placement:100  version:3  modified_at_second:1614937885 
          volume id:317  size:1076215096  collection:"collection2"  file_count:23572  delete_count:2  deleted_byte_count:1441356  replica_placement:100  version:3  modified_at_second:1614943965 
          volume id:318  size:1075965168  collection:"collection2"  file_count:22459  delete_count:2  deleted_byte_count:37778  replica_placement:100  version:3  modified_at_second:1614943862 
          volume id:319  size:1073952880  collection:"collection2"  file_count:22286  delete_count:2  deleted_byte_count:43421  replica_placement:100  version:3  modified_at_second:1614943810 
          volume id:320  size:1082437792  collection:"collection2"  file_count:21544  delete_count:3  deleted_byte_count:16712  replica_placement:100  version:3  modified_at_second:1614943599 
          volume id:321  size:1081477904  collection:"collection2"  file_count:23531  delete_count:5  deleted_byte_count:262564  replica_placement:100  version:3  modified_at_second:1614943982 
          volume id:324  size:1075606680  collection:"collection2"  file_count:20799  delete_count:1  deleted_byte_count:251210  replica_placement:100  version:3  modified_at_second:1614950310 
          volume id:325  size:1080701144  collection:"collection2"  file_count:21735  replica_placement:100  version:3  modified_at_second:1614950525 
          volume id:330  size:1080825832  collection:"collection2"  file_count:22464  delete_count:2  deleted_byte_count:15771  replica_placement:100  version:3  modified_at_second:1614956477 
          volume id:332  size:1075569928  collection:"collection2"  file_count:22097  delete_count:3  deleted_byte_count:98273  replica_placement:100  version:3  modified_at_second:1614962869 
          volume id:334  size:1075607880  collection:"collection2"  file_count:22546  delete_count:6  deleted_byte_count:101538  replica_placement:100  version:3  modified_at_second:1614962978 
          volume id:336  size:1087853056  collection:"collection2"  file_count:22801  delete_count:2  deleted_byte_count:26394  replica_placement:100  version:3  modified_at_second:1614963005 
          volume id:337  size:1075646784  collection:"collection2"  file_count:21934  delete_count:1  deleted_byte_count:3397  replica_placement:100  version:3  modified_at_second:1614969937 
          volume id:338  size:1076118304  collection:"collection2"  file_count:21680  replica_placement:100  version:3  modified_at_second:1614969850 
          volume id:340  size:1079462184  collection:"collection2"  file_count:22319  delete_count:4  deleted_byte_count:93620  replica_placement:100  version:3  modified_at_second:1614969978 
          volume id:341  size:1074448400  collection:"collection2"  file_count:21590  delete_count:5  deleted_byte_count:160085  replica_placement:100  version:3  modified_at_second:1614969858 
          volume id:342  size:1080186424  collection:"collection2"  file_count:22405  delete_count:4  deleted_byte_count:64819  replica_placement:100  version:3  modified_at_second:1614977521 
          volume id:344  size:1075035416  collection:"collection2"  file_count:21765  delete_count:1  deleted_byte_count:24623  replica_placement:100  version:3  modified_at_second:1614977465 
          volume id:345  size:1074560760  collection:"collection2"  file_count:22117  delete_count:2  deleted_byte_count:373286  replica_placement:100  version:3  modified_at_second:1614977457 
          volume id:346  size:1076464112  collection:"collection2"  file_count:22320  delete_count:4  deleted_byte_count:798258  replica_placement:100  version:3  modified_at_second:1615631322 
          volume id:348  size:1080623640  collection:"collection2"  file_count:21667  delete_count:1  deleted_byte_count:2443  replica_placement:100  version:3  modified_at_second:1614984606 
          volume id:350  size:1074756688  collection:"collection2"  file_count:21990  delete_count:3  deleted_byte_count:233881  replica_placement:100  version:3  modified_at_second:1614984682 
          volume id:351  size:1078795112  collection:"collection2"  file_count:23660  delete_count:3  deleted_byte_count:102141  replica_placement:100  version:3  modified_at_second:1614984816 
          volume id:352  size:1077145936  collection:"collection2"  file_count:22066  delete_count:1  deleted_byte_count:1018  replica_placement:100  version:3  modified_at_second:1614992130 
          volume id:353  size:1074897496  collection:"collection2"  file_count:21266  delete_count:2  deleted_byte_count:3105374  replica_placement:100  version:3  modified_at_second:1614991951 
          volume id:355  size:1075239728  collection:"collection2"  file_count:22244  delete_count:1  deleted_byte_count:23282  replica_placement:100  version:3  modified_at_second:1614992157 
          volume id:356  size:1083305048  collection:"collection2"  file_count:21552  delete_count:4  deleted_byte_count:14472  replica_placement:100  version:3  modified_at_second:1614992028 
          volume id:358  size:1085152368  collection:"collection2"  file_count:23756  delete_count:3  deleted_byte_count:44531  replica_placement:100  version:3  modified_at_second:1614998824 
          volume id:360  size:1075532456  collection:"collection2"  file_count:22574  delete_count:3  deleted_byte_count:1774776  replica_placement:100  version:3  modified_at_second:1614998770 
          volume id:361  size:1075362744  collection:"collection2"  file_count:22272  delete_count:1  deleted_byte_count:3497  replica_placement:100  version:3  modified_at_second:1614998668 
          volume id:375  size:1076140568  collection:"collection2"  file_count:21880  delete_count:2  deleted_byte_count:51103  replica_placement:100  version:3  modified_at_second:1615016787 
          volume id:376  size:1074845944  collection:"collection2"  file_count:22908  delete_count:4  deleted_byte_count:432305  replica_placement:100  version:3  modified_at_second:1615016916 
          volume id:377  size:957284144  collection:"collection2"  file_count:14923  delete_count:1  deleted_byte_count:1797  replica_placement:100  version:3  modified_at_second:1615632323 
          volume id:378  size:959273936  collection:"collection2"  file_count:15027  delete_count:4  deleted_byte_count:231414  replica_placement:100  version:3  modified_at_second:1615632323 
          volume id:381  size:939261032  collection:"collection2"  file_count:14615  delete_count:5  deleted_byte_count:1192272  replica_placement:100  version:3  modified_at_second:1615632324 
        Disk hdd total size:111617646696 file_count:1762773 deleted_file:1221 deleted_bytes:398484585 
      DataNode 192.168.1.6:8080 total size:111617646696 file_count:1762773 deleted_file:1221 deleted_bytes:398484585 
    Rack rack3 total size:111617646696 file_count:1762773 deleted_file:1221 deleted_bytes:398484585 
  DataCenter dc3 total size:111617646696 file_count:1762773 deleted_file:1221 deleted_bytes:398484585 
  DataCenter dc4 hdd(volume:267/2000 active:267 free:1733 remote:0)
    Rack DefaultRack hdd(volume:267/2000 active:267 free:1733 remote:0)
      DataNode 192.168.1.1:8080 hdd(volume:267/2000 active:267 free:1733 remote:0)
        Disk hdd(volume:267/2000 active:267 free:1733 remote:0)
          volume id:1  size:284693256  file_count:1558  delete_count:2  deleted_byte_count:4818  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632062 
          volume id:2  size:289228560  file_count:1640  delete_count:4  deleted_byte_count:464508  replica_placement:100  version:3  compact_revision:6  modified_at_second:1615630622 
          volume id:3  size:308741952  file_count:1637  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632242 
          volume id:4  size:285986968  file_count:1640  delete_count:1  deleted_byte_count:145095  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632302 
          volume id:5  size:293806008  file_count:1669  delete_count:2  deleted_byte_count:274334  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631342 
          volume id:6  size:302411024  file_count:1604  delete_count:2  deleted_byte_count:274587  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631402 
          volume id:7  size:1924728  collection:"collection4"  file_count:15  replica_placement:100  version:3  modified_at_second:1609331040 
          volume id:9  size:77337416  collection:"collection3"  file_count:58  replica_placement:100  version:3  ttl:772  modified_at_second:1615513762 
          volume id:10  size:1212784656  collection:"collection0"  file_count:58  replica_placement:100  version:3  modified_at_second:1609814543 
          volume id:11  size:1109224552  collection:"collection0"  file_count:44  replica_placement:100  version:3  modified_at_second:1609815123 
          volume id:12  size:1110923848  collection:"collection0"  file_count:45  replica_placement:100  version:3  modified_at_second:1609819726 
          volume id:13  size:1184910656  collection:"collection0"  file_count:47  replica_placement:100  version:3  modified_at_second:1609827832 
          volume id:14  size:1107475720  collection:"collection0"  file_count:80  delete_count:3  deleted_byte_count:6870  replica_placement:100  version:3  modified_at_second:1612956983 
          volume id:15  size:1115965160  collection:"collection0"  file_count:83  delete_count:3  deleted_byte_count:4956  replica_placement:100  version:3  modified_at_second:1612957001 
          volume id:16  size:1113666048  collection:"collection0"  file_count:73  delete_count:5  deleted_byte_count:6318  replica_placement:100  version:3  modified_at_second:1612957007 
          volume id:17  size:1095115800  collection:"collection0"  file_count:83  delete_count:3  deleted_byte_count:7099  replica_placement:100  version:3  modified_at_second:1612957000 
          volume id:18  size:1096678688  collection:"collection0"  file_count:88  delete_count:4  deleted_byte_count:8633  replica_placement:100  version:3  modified_at_second:1612957000 
          volume id:19  size:1096923792  collection:"collection0"  file_count:100  delete_count:10  deleted_byte_count:75779917  replica_placement:100  version:3  compact_revision:4  modified_at_second:1612957011 
          volume id:20  size:1074760432  collection:"collection0"  file_count:82  delete_count:5  deleted_byte_count:12156  replica_placement:100  version:3  compact_revision:2  modified_at_second:1612957011 
          volume id:22  size:1086828368  collection:"collection0"  file_count:75  delete_count:3  deleted_byte_count:5551  replica_placement:100  version:3  modified_at_second:1612957007 
          volume id:23  size:1076380280  collection:"collection0"  file_count:68  delete_count:2  deleted_byte_count:2910  replica_placement:100  version:3  modified_at_second:1612957011 
          volume id:24  size:1074139808  collection:"collection0"  file_count:90  delete_count:1  deleted_byte_count:1977  replica_placement:100  version:3  modified_at_second:1612957011 
          volume id:25  size:690757544  collection:"collection0"  file_count:38  delete_count:1  deleted_byte_count:1944  replica_placement:100  version:3  modified_at_second:1612956995 
          volume id:26  size:532657632  collection:"collection0"  file_count:100  delete_count:4  deleted_byte_count:9081  replica_placement:100  version:3  modified_at_second:1614170023 
          volume id:34  size:1077111136  collection:"collection1"  file_count:9781  delete_count:110  deleted_byte_count:20894827  replica_placement:100  version:3  modified_at_second:1615619366 
          volume id:35  size:1075241656  collection:"collection1"  file_count:10523  delete_count:96  deleted_byte_count:46618989  replica_placement:100  version:3  modified_at_second:1615618790 
          volume id:36  size:1075118360  collection:"collection1"  file_count:10342  delete_count:116  deleted_byte_count:25493106  replica_placement:100  version:3  modified_at_second:1615606148 
          volume id:37  size:1075895584  collection:"collection1"  file_count:12013  delete_count:98  deleted_byte_count:50747932  replica_placement:100  version:3  modified_at_second:1615594777 
          volume id:39  size:1076606536  collection:"collection1"  file_count:12612  delete_count:78  deleted_byte_count:17462730  replica_placement:100  version:3  modified_at_second:1615611959 
          volume id:40  size:1075358552  collection:"collection1"  file_count:12597  delete_count:62  deleted_byte_count:11657901  replica_placement:100  version:3  modified_at_second:1615612994 
          volume id:41  size:1076283528  collection:"collection1"  file_count:12088  delete_count:84  deleted_byte_count:19319268  replica_placement:100  version:3  modified_at_second:1615596736 
          volume id:42  size:1093948352  collection:"collection1"  file_count:7889  delete_count:47  deleted_byte_count:5697275  replica_placement:100  version:3  modified_at_second:1615548908 
          volume id:43  size:1116445864  collection:"collection1"  file_count:7358  delete_count:54  deleted_byte_count:9534379  replica_placement:100  version:3  modified_at_second:1615566170 
          volume id:44  size:1077582560  collection:"collection1"  file_count:7295  delete_count:50  deleted_byte_count:12618414  replica_placement:100  version:3  modified_at_second:1615566170 
          volume id:45  size:1075254640  collection:"collection1"  file_count:10772  delete_count:76  deleted_byte_count:22426345  replica_placement:100  version:3  modified_at_second:1615573499 
          volume id:46  size:1075286056  collection:"collection1"  file_count:9947  delete_count:309  deleted_byte_count:105601163  replica_placement:100  version:3  modified_at_second:1615569826 
          volume id:48  size:1076778720  collection:"collection1"  file_count:9850  delete_count:77  deleted_byte_count:16641907  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630690 
          volume id:50  size:1076688224  collection:"collection1"  file_count:7921  delete_count:26  deleted_byte_count:5162032  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615610879 
          volume id:52  size:1083529704  collection:"collection1"  file_count:10128  delete_count:32  deleted_byte_count:10608391  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615599195 
          volume id:53  size:1063089216  collection:"collection1"  file_count:9832  delete_count:31  deleted_byte_count:9273066  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632156 
          volume id:55  size:1012890016  collection:"collection1"  file_count:8651  delete_count:27  deleted_byte_count:9418841  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631452 
          volume id:57  size:839849792  collection:"collection1"  file_count:7514  delete_count:24  deleted_byte_count:6228543  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631774 
          volume id:58  size:908064200  collection:"collection1"  file_count:8128  delete_count:21  deleted_byte_count:6113731  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632342 
          volume id:59  size:988302272  collection:"collection1"  file_count:8098  delete_count:20  deleted_byte_count:3947615  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632238 
          volume id:60  size:1010702480  collection:"collection1"  file_count:8969  delete_count:79  deleted_byte_count:24782814  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615632439 
          volume id:61  size:975604488  collection:"collection1"  file_count:8683  delete_count:20  deleted_byte_count:10276072  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615631176 
          volume id:62  size:873845936  collection:"collection1"  file_count:7897  delete_count:23  deleted_byte_count:10920170  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631133 
          volume id:64  size:965638488  collection:"collection1"  file_count:8218  delete_count:27  deleted_byte_count:6922489  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631031 
          volume id:65  size:823283552  collection:"collection1"  file_count:7834  delete_count:29  deleted_byte_count:5950610  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632306 
          volume id:66  size:821343440  collection:"collection1"  file_count:7383  delete_count:29  deleted_byte_count:12010343  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631968 
          volume id:67  size:878713872  collection:"collection1"  file_count:7299  delete_count:117  deleted_byte_count:24857326  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632156 
          volume id:68  size:898630584  collection:"collection1"  file_count:6934  delete_count:95  deleted_byte_count:27460707  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632284 
          volume id:70  size:886695472  collection:"collection1"  file_count:7769  delete_count:164  deleted_byte_count:45162513  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632398 
          volume id:71  size:907608392  collection:"collection1"  file_count:7658  delete_count:122  deleted_byte_count:27622941  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632307 
          volume id:72  size:903990720  collection:"collection1"  file_count:6996  delete_count:240  deleted_byte_count:74147727  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630982 
          volume id:73  size:929047664  collection:"collection1"  file_count:7038  delete_count:227  deleted_byte_count:65336664  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630707 
          volume id:74  size:957046128  collection:"collection1"  file_count:6981  delete_count:259  deleted_byte_count:73080838  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615631460 
          volume id:75  size:908044992  collection:"collection1"  file_count:6911  delete_count:268  deleted_byte_count:73934373  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632430 
          volume id:76  size:985296744  collection:"collection1"  file_count:6566  delete_count:61  deleted_byte_count:44464430  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632284 
          volume id:77  size:929398296  collection:"collection1"  file_count:7427  delete_count:238  deleted_byte_count:59581579  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632013 
          volume id:78  size:1075671512  collection:"collection1"  file_count:7540  delete_count:258  deleted_byte_count:71726846  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615582829 
          volume id:79  size:948225472  collection:"collection1"  file_count:6997  delete_count:227  deleted_byte_count:60625763  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631326 
          volume id:82  size:1041661800  collection:"collection1"  file_count:7043  delete_count:207  deleted_byte_count:52275724  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632430 
          volume id:83  size:936195856  collection:"collection1"  file_count:7593  delete_count:13  deleted_byte_count:4633917  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632029 
          volume id:85  size:1023867520  collection:"collection1"  file_count:7787  delete_count:240  deleted_byte_count:82091742  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631723 
          volume id:86  size:1009437488  collection:"collection1"  file_count:8474  delete_count:236  deleted_byte_count:64543674  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630812 
          volume id:87  size:922276640  collection:"collection1"  file_count:12902  delete_count:13  deleted_byte_count:3412959  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632438 
          volume id:89  size:1044401976  collection:"collection1"  file_count:14943  delete_count:243  deleted_byte_count:58543159  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632208 
          volume id:90  size:891145784  collection:"collection1"  file_count:14608  delete_count:10  deleted_byte_count:2564369  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615629390 
          volume id:91  size:936572832  collection:"collection1"  file_count:14686  delete_count:11  deleted_byte_count:4717727  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631851 
          volume id:92  size:992440712  collection:"collection1"  file_count:7061  delete_count:195  deleted_byte_count:60649573  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630566 
          volume id:93  size:1079603768  collection:"collection1"  file_count:7878  delete_count:270  deleted_byte_count:74150048  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615556015 
          volume id:94  size:1030685824  collection:"collection1"  file_count:7660  delete_count:207  deleted_byte_count:70150733  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631616 
          volume id:95  size:990879168  collection:"collection1"  file_count:6620  delete_count:206  deleted_byte_count:60363604  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615631866 
          volume id:96  size:989296136  collection:"collection1"  file_count:7544  delete_count:229  deleted_byte_count:59931853  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630778 
          volume id:97  size:1053112992  collection:"collection1"  file_count:6789  delete_count:50  deleted_byte_count:38894001  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631194 
          volume id:99  size:1071718504  collection:"collection1"  file_count:7470  delete_count:8  deleted_byte_count:9624950  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631175 
          volume id:100  size:1083617440  collection:"collection1"  file_count:7018  delete_count:187  deleted_byte_count:61304236  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615505917 
          volume id:101  size:1077109520  collection:"collection1"  file_count:7706  delete_count:226  deleted_byte_count:77864841  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630994 
          volume id:102  size:1074359920  collection:"collection1"  file_count:7338  delete_count:7  deleted_byte_count:6499151  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615626683 
          volume id:103  size:1075863904  collection:"collection1"  file_count:7184  delete_count:186  deleted_byte_count:58872238  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615628417 
          volume id:104  size:1076383768  collection:"collection1"  file_count:7663  delete_count:184  deleted_byte_count:100578087  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615602661 
          volume id:105  size:1073996824  collection:"collection1"  file_count:6873  delete_count:19  deleted_byte_count:14271533  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615499756 
          volume id:108  size:1074648024  collection:"collection1"  file_count:7472  delete_count:194  deleted_byte_count:70864699  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615593232 
          volume id:109  size:1075254560  collection:"collection1"  file_count:7556  delete_count:263  deleted_byte_count:55155265  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615502487 
          volume id:110  size:1076575744  collection:"collection1"  file_count:6996  delete_count:163  deleted_byte_count:52954032  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615590786 
          volume id:111  size:1073826232  collection:"collection1"  file_count:7355  delete_count:155  deleted_byte_count:50083578  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615593233 
          volume id:114  size:1074762784  collection:"collection1"  file_count:8802  delete_count:156  deleted_byte_count:38470055  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615591826 
          volume id:115  size:1076192240  collection:"collection1"  file_count:7690  delete_count:154  deleted_byte_count:32267193  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615285295 
          volume id:116  size:1074489504  collection:"collection1"  file_count:9981  delete_count:174  deleted_byte_count:53998777  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615611567 
          volume id:117  size:1073917192  collection:"collection1"  file_count:9520  delete_count:114  deleted_byte_count:21835126  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615573714 
          volume id:118  size:1074064400  collection:"collection1"  file_count:8738  delete_count:15  deleted_byte_count:3460697  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615516265 
          volume id:119  size:1075940104  collection:"collection1"  file_count:9003  delete_count:12  deleted_byte_count:9128155  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615573880 
          volume id:120  size:1076115928  collection:"collection1"  file_count:9639  delete_count:118  deleted_byte_count:33357871  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615482567 
          volume id:121  size:1078803248  collection:"collection1"  file_count:10113  delete_count:441  deleted_byte_count:94128627  replica_placement:100  version:3  modified_at_second:1615506629 
          volume id:122  size:1076235312  collection:"collection1"  file_count:9106  delete_count:252  deleted_byte_count:93041272  replica_placement:100  version:3  modified_at_second:1615585913 
          volume id:123  size:1080491112  collection:"collection1"  file_count:10623  delete_count:302  deleted_byte_count:83956998  replica_placement:100  version:3  modified_at_second:1615585916 
          volume id:124  size:1074519360  collection:"collection1"  file_count:9457  delete_count:286  deleted_byte_count:74752459  replica_placement:100  version:3  modified_at_second:1615585916 
          volume id:125  size:1088687040  collection:"collection1"  file_count:9518  delete_count:281  deleted_byte_count:76037905  replica_placement:100  version:3  modified_at_second:1615585913 
          volume id:126  size:1073867464  collection:"collection1"  file_count:9320  delete_count:278  deleted_byte_count:94547424  replica_placement:100  version:3  modified_at_second:1615585912 
          volume id:127  size:1074907336  collection:"collection1"  file_count:9900  delete_count:133  deleted_byte_count:48570820  replica_placement:100  version:3  modified_at_second:1615612991 
          volume id:129  size:1074704272  collection:"collection1"  file_count:10012  delete_count:150  deleted_byte_count:64491721  replica_placement:100  version:3  modified_at_second:1615627566 
          volume id:130  size:1075000632  collection:"collection1"  file_count:10633  delete_count:161  deleted_byte_count:34768201  replica_placement:100  version:3  modified_at_second:1615582330 
          volume id:131  size:1075279584  collection:"collection1"  file_count:10075  delete_count:135  deleted_byte_count:29795712  replica_placement:100  version:3  modified_at_second:1615523898 
          volume id:132  size:1088539496  collection:"collection1"  file_count:11051  delete_count:71  deleted_byte_count:17178322  replica_placement:100  version:3  modified_at_second:1615619584 
          volume id:133  size:1075952760  collection:"collection1"  file_count:9538  delete_count:74  deleted_byte_count:19558008  replica_placement:100  version:3  modified_at_second:1615584780 
          volume id:134  size:1074367304  collection:"collection1"  file_count:10662  delete_count:69  deleted_byte_count:25530139  replica_placement:100  version:3  modified_at_second:1615555876 
          volume id:135  size:1073906720  collection:"collection1"  file_count:10446  delete_count:71  deleted_byte_count:28599975  replica_placement:100  version:3  modified_at_second:1615569816 
          volume id:137  size:1074309264  collection:"collection1"  file_count:9633  delete_count:50  deleted_byte_count:27487972  replica_placement:100  version:3  modified_at_second:1615572231 
          volume id:139  size:1074163936  collection:"collection1"  file_count:9314  delete_count:43  deleted_byte_count:10631353  replica_placement:100  version:3  modified_at_second:1615571946 
          volume id:141  size:1074619488  collection:"collection1"  file_count:9840  delete_count:45  deleted_byte_count:40890181  replica_placement:100  version:3  modified_at_second:1615630994 
          volume id:142  size:1075732992  collection:"collection1"  file_count:9009  delete_count:48  deleted_byte_count:9912854  replica_placement:100  version:3  modified_at_second:1615598914 
          volume id:143  size:1075011280  collection:"collection1"  file_count:9608  delete_count:51  deleted_byte_count:37282460  replica_placement:100  version:3  modified_at_second:1615488586 
          volume id:145  size:1074394928  collection:"collection1"  file_count:9255  delete_count:34  deleted_byte_count:38011392  replica_placement:100  version:3  modified_at_second:1615591825 
          volume id:146  size:1076337520  collection:"collection1"  file_count:10492  delete_count:50  deleted_byte_count:17071505  replica_placement:100  version:3  modified_at_second:1615632005 
          volume id:147  size:1077130544  collection:"collection1"  file_count:10451  delete_count:27  deleted_byte_count:8290907  replica_placement:100  version:3  modified_at_second:1615604117 
          volume id:148  size:1076066568  collection:"collection1"  file_count:9547  delete_count:33  deleted_byte_count:7034089  replica_placement:100  version:3  modified_at_second:1615586393 
          volume id:149  size:1074989016  collection:"collection1"  file_count:8352  delete_count:35  deleted_byte_count:7179742  replica_placement:100  version:3  modified_at_second:1615494496 
          volume id:150  size:1076290408  collection:"collection1"  file_count:9328  delete_count:33  deleted_byte_count:43417791  replica_placement:100  version:3  modified_at_second:1615611569 
          volume id:151  size:1098659752  collection:"collection1"  file_count:10805  delete_count:27  deleted_byte_count:7209106  replica_placement:100  version:3  modified_at_second:1615586390 
          volume id:152  size:1075941376  collection:"collection1"  file_count:9951  delete_count:36  deleted_byte_count:25348335  replica_placement:100  version:3  modified_at_second:1615606614 
          volume id:153  size:1078539784  collection:"collection1"  file_count:10924  delete_count:34  deleted_byte_count:12603081  replica_placement:100  version:3  modified_at_second:1615606614 
          volume id:154  size:1081244752  collection:"collection1"  file_count:11002  delete_count:31  deleted_byte_count:8467560  replica_placement:100  version:3  modified_at_second:1615478471 
          volume id:156  size:1074975832  collection:"collection1"  file_count:9535  delete_count:40  deleted_byte_count:11426621  replica_placement:100  version:3  modified_at_second:1615628342 
          volume id:157  size:1076758536  collection:"collection1"  file_count:10012  delete_count:19  deleted_byte_count:11688737  replica_placement:100  version:3  modified_at_second:1615597782 
          volume id:158  size:1087251976  collection:"collection1"  file_count:9972  delete_count:20  deleted_byte_count:10328429  replica_placement:100  version:3  modified_at_second:1615588879 
          volume id:159  size:1074132336  collection:"collection1"  file_count:9382  delete_count:27  deleted_byte_count:11474574  replica_placement:100  version:3  modified_at_second:1615593593 
          volume id:160  size:1075680976  collection:"collection1"  file_count:9772  delete_count:22  deleted_byte_count:4981968  replica_placement:100  version:3  modified_at_second:1615597782 
          volume id:161  size:1077397136  collection:"collection1"  file_count:9988  delete_count:28  deleted_byte_count:12509164  replica_placement:100  version:3  modified_at_second:1615631452 
          volume id:162  size:1074286880  collection:"collection1"  file_count:11220  delete_count:17  deleted_byte_count:1815547  replica_placement:100  version:3  modified_at_second:1615478127 
          volume id:163  size:1074457224  collection:"collection1"  file_count:12524  delete_count:27  deleted_byte_count:6359619  replica_placement:100  version:3  modified_at_second:1615579313 
          volume id:164  size:1074261256  collection:"collection1"  file_count:11922  delete_count:25  deleted_byte_count:2923288  replica_placement:100  version:3  modified_at_second:1615620085 
          volume id:165  size:1073891080  collection:"collection1"  file_count:9152  delete_count:12  deleted_byte_count:19164659  replica_placement:100  version:3  modified_at_second:1615471910 
          volume id:166  size:1075637536  collection:"collection1"  file_count:14211  delete_count:24  deleted_byte_count:20415490  replica_placement:100  version:3  modified_at_second:1615491021 
          volume id:167  size:1073958280  collection:"collection1"  file_count:25231  delete_count:48  deleted_byte_count:26022344  replica_placement:100  version:3  modified_at_second:1615620014 
          volume id:168  size:1074718864  collection:"collection1"  file_count:25702  delete_count:40  deleted_byte_count:4024775  replica_placement:100  version:3  modified_at_second:1615585664 
          volume id:169  size:1073863032  collection:"collection1"  file_count:25248  delete_count:43  deleted_byte_count:3013817  replica_placement:100  version:3  modified_at_second:1615569832 
          volume id:170  size:1075747088  collection:"collection1"  file_count:24596  delete_count:41  deleted_byte_count:3494711  replica_placement:100  version:3  modified_at_second:1615579207 
          volume id:171  size:1081881400  collection:"collection1"  file_count:24215  delete_count:36  deleted_byte_count:3191335  replica_placement:100  version:3  modified_at_second:1615596486 
          volume id:172  size:1074787304  collection:"collection1"  file_count:31236  delete_count:50  deleted_byte_count:3316482  replica_placement:100  version:3  modified_at_second:1615612385 
          volume id:174  size:1073824160  collection:"collection1"  file_count:30689  delete_count:36  deleted_byte_count:2160116  replica_placement:100  version:3  modified_at_second:1615598914 
          volume id:175  size:1077742472  collection:"collection1"  file_count:32353  delete_count:33  deleted_byte_count:1861403  replica_placement:100  version:3  modified_at_second:1615559516 
          volume id:176  size:1073854800  collection:"collection1"  file_count:30582  delete_count:34  deleted_byte_count:7701976  replica_placement:100  version:3  modified_at_second:1615626169 
          volume id:178  size:1087560112  collection:"collection1"  file_count:23482  delete_count:22  deleted_byte_count:18810492  replica_placement:100  version:3  modified_at_second:1615541369 
          volume id:179  size:1074313920  collection:"collection1"  file_count:21829  delete_count:24  deleted_byte_count:45574435  replica_placement:100  version:3  modified_at_second:1615580308 
          volume id:180  size:1078438448  collection:"collection1"  file_count:23614  delete_count:12  deleted_byte_count:4496474  replica_placement:100  version:3  modified_at_second:1614773243 
          volume id:181  size:1074571672  collection:"collection1"  file_count:22898  delete_count:19  deleted_byte_count:6628413  replica_placement:100  version:3  modified_at_second:1614745117 
          volume id:183  size:1076361616  collection:"collection1"  file_count:31293  delete_count:16  deleted_byte_count:468841  replica_placement:100  version:3  modified_at_second:1615572985 
          volume id:184  size:1074594216  collection:"collection1"  file_count:31368  delete_count:22  deleted_byte_count:857453  replica_placement:100  version:3  modified_at_second:1615586578 
          volume id:185  size:1074099592  collection:"collection1"  file_count:30612  delete_count:17  deleted_byte_count:2610847  replica_placement:100  version:3  modified_at_second:1615506835 
          volume id:186  size:1074220664  collection:"collection1"  file_count:31450  delete_count:15  deleted_byte_count:391855  replica_placement:100  version:3  modified_at_second:1615614934 
          volume id:187  size:1074396112  collection:"collection1"  file_count:31853  delete_count:17  deleted_byte_count:454283  replica_placement:100  version:3  modified_at_second:1615590491 
          volume id:188  size:1074732632  collection:"collection1"  file_count:31867  delete_count:19  deleted_byte_count:393743  replica_placement:100  version:3  modified_at_second:1615487645 
          volume id:189  size:1074847824  collection:"collection1"  file_count:31450  delete_count:16  deleted_byte_count:1040552  replica_placement:100  version:3  modified_at_second:1615335661 
          volume id:190  size:1074008968  collection:"collection1"  file_count:31987  delete_count:11  deleted_byte_count:685125  replica_placement:100  version:3  modified_at_second:1615447162 
          volume id:191  size:1075492960  collection:"collection1"  file_count:31301  delete_count:19  deleted_byte_count:708401  replica_placement:100  version:3  modified_at_second:1615357457 
          volume id:192  size:1075857384  collection:"collection1"  file_count:31490  delete_count:25  deleted_byte_count:720617  replica_placement:100  version:3  modified_at_second:1615621632 
          volume id:193  size:1076616760  collection:"collection1"  file_count:31907  delete_count:16  deleted_byte_count:464900  replica_placement:100  version:3  modified_at_second:1615507877 
          volume id:194  size:1073985792  collection:"collection1"  file_count:31434  delete_count:18  deleted_byte_count:391432  replica_placement:100  version:3  modified_at_second:1615559502 
          volume id:195  size:1074158304  collection:"collection1"  file_count:31453  delete_count:15  deleted_byte_count:718266  replica_placement:100  version:3  modified_at_second:1615559331 
          volume id:196  size:1074594640  collection:"collection1"  file_count:31665  delete_count:18  deleted_byte_count:3468922  replica_placement:100  version:3  modified_at_second:1615501690 
          volume id:198  size:1075104624  collection:"collection1"  file_count:16577  delete_count:18  deleted_byte_count:6583181  replica_placement:100  version:3  modified_at_second:1615623371 
          volume id:199  size:1078117688  collection:"collection1"  file_count:16497  delete_count:14  deleted_byte_count:1514286  replica_placement:100  version:3  modified_at_second:1615585987 
          volume id:200  size:1075630464  collection:"collection1"  file_count:16380  delete_count:18  deleted_byte_count:1103109  replica_placement:100  version:3  modified_at_second:1615485252 
          volume id:201  size:1091460440  collection:"collection1"  file_count:16684  delete_count:26  deleted_byte_count:5590335  replica_placement:100  version:3  modified_at_second:1615585987 
          volume id:204  size:1079766904  collection:"collection1"  file_count:3233  delete_count:255  deleted_byte_count:104707641  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615565702 
          volume id:207  size:1081939960  collection:"collection1"  file_count:3010  delete_count:4  deleted_byte_count:692350  replica_placement:100  version:3  modified_at_second:1615269061 
          volume id:208  size:1077863624  collection:"collection1"  file_count:3147  delete_count:6  deleted_byte_count:858726  replica_placement:100  version:3  modified_at_second:1615495515 
          volume id:209  size:1074083592  collection:"collection1"  file_count:3238  delete_count:4  deleted_byte_count:1494244  replica_placement:100  version:3  modified_at_second:1615419954 
          volume id:210  size:1094311304  collection:"collection1"  file_count:3468  delete_count:4  deleted_byte_count:466433  replica_placement:100  version:3  modified_at_second:1615495515 
          volume id:211  size:1080610712  collection:"collection1"  file_count:3247  delete_count:7  deleted_byte_count:1891456  replica_placement:100  version:3  modified_at_second:1615269124 
          volume id:216  size:1080073496  collection:"collection1"  file_count:3316  delete_count:4  deleted_byte_count:179819  replica_placement:100  version:3  modified_at_second:1615586387 
          volume id:218  size:1081263944  collection:"collection1"  file_count:3433  delete_count:14  deleted_byte_count:3454237  replica_placement:100  version:3  modified_at_second:1615603637 
          volume id:220  size:1081928312  collection:"collection1"  file_count:3166  delete_count:13  deleted_byte_count:4127709  replica_placement:100  version:3  modified_at_second:1615579317 
          volume id:221  size:1106545536  collection:"collection1"  file_count:3153  delete_count:11  deleted_byte_count:1496835  replica_placement:100  version:3  modified_at_second:1615269138 
          volume id:224  size:1093691520  collection:"collection1"  file_count:3463  delete_count:10  deleted_byte_count:1128328  replica_placement:100  version:3  modified_at_second:1615601870 
          volume id:225  size:1080698928  collection:"collection1"  file_count:3115  delete_count:7  deleted_byte_count:18170416  replica_placement:100  version:3  modified_at_second:1615434685 
          volume id:226  size:1103504792  collection:"collection1"  file_count:2965  delete_count:10  deleted_byte_count:2639254  replica_placement:100  version:3  modified_at_second:1615601870 
          volume id:227  size:1106699864  collection:"collection1"  file_count:2827  delete_count:19  deleted_byte_count:5393310  replica_placement:100  version:3  modified_at_second:1615609989 
          volume id:228  size:1109784072  collection:"collection1"  file_count:2504  delete_count:24  deleted_byte_count:5458950  replica_placement:100  version:3  modified_at_second:1615610489 
          volume id:229  size:1109855256  collection:"collection1"  file_count:2857  delete_count:22  deleted_byte_count:2839883  replica_placement:100  version:3  modified_at_second:1615609989 
          volume id:231  size:1112917664  collection:"collection1"  file_count:3151  delete_count:19  deleted_byte_count:2852517  replica_placement:100  version:3  modified_at_second:1615611350 
          volume id:232  size:1073901520  collection:"collection1"  file_count:3004  delete_count:54  deleted_byte_count:10273081  replica_placement:100  version:3  modified_at_second:1615611352 
          volume id:233  size:1080526464  collection:"collection1"  file_count:3136  delete_count:61  deleted_byte_count:17991717  replica_placement:100  version:3  modified_at_second:1615611354 
          volume id:236  size:1089476200  collection:"collection1"  file_count:3231  delete_count:53  deleted_byte_count:11625921  replica_placement:100  version:3  modified_at_second:1615611351 
          volume id:238  size:354320000  collection:"collection1"  file_count:701  delete_count:17  deleted_byte_count:5940420  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615632030 
          volume id:240  size:424791528  collection:"collection1"  file_count:734  delete_count:12  deleted_byte_count:7353071  replica_placement:100  version:3  modified_at_second:1615631669 
          volume id:242  size:1075383304  collection:"collection2"  file_count:10470  replica_placement:100  version:3  modified_at_second:1614852115 
          volume id:243  size:1088174560  collection:"collection2"  file_count:11109  delete_count:1  deleted_byte_count:938  replica_placement:100  version:3  modified_at_second:1614852202 
          volume id:245  size:1074597056  collection:"collection2"  file_count:10371  delete_count:3  deleted_byte_count:209701  replica_placement:100  version:3  modified_at_second:1614852093 
          volume id:247  size:1075859784  collection:"collection2"  file_count:10443  delete_count:2  deleted_byte_count:564486  replica_placement:100  version:3  modified_at_second:1614856152 
          volume id:249  size:1074819168  collection:"collection2"  file_count:10763  delete_count:2  deleted_byte_count:271699  replica_placement:100  version:3  modified_at_second:1614856231 
          volume id:250  size:1080572256  collection:"collection2"  file_count:10220  replica_placement:100  version:3  modified_at_second:1614856129 
          volume id:251  size:1075684408  collection:"collection2"  file_count:10847  replica_placement:100  version:3  modified_at_second:1614856270 
          volume id:254  size:1074830800  collection:"collection2"  file_count:14140  delete_count:2  deleted_byte_count:105892  replica_placement:100  version:3  modified_at_second:1614861115 
          volume id:257  size:1082621664  collection:"collection2"  file_count:18172  delete_count:2  deleted_byte_count:25125  replica_placement:100  version:3  modified_at_second:1614866395 
          volume id:260  size:1075105664  collection:"collection2"  file_count:17316  delete_count:4  deleted_byte_count:2015310  replica_placement:100  version:3  modified_at_second:1614866226 
          volume id:261  size:1076628592  collection:"collection2"  file_count:18355  delete_count:1  deleted_byte_count:1155  replica_placement:100  version:3  modified_at_second:1614866420 
          volume id:262  size:1078492464  collection:"collection2"  file_count:20390  delete_count:3  deleted_byte_count:287601  replica_placement:100  version:3  modified_at_second:1614871601 
          volume id:263  size:1077167440  collection:"collection2"  file_count:20227  delete_count:4  deleted_byte_count:97887  replica_placement:100  version:3  modified_at_second:1614871567 
          volume id:268  size:1074490592  collection:"collection2"  file_count:21698  delete_count:1  deleted_byte_count:33968  replica_placement:100  version:3  modified_at_second:1614877435 
          volume id:269  size:1077552720  collection:"collection2"  file_count:21875  delete_count:4  deleted_byte_count:347272  replica_placement:100  version:3  modified_at_second:1614877481 
          volume id:271  size:1076992648  collection:"collection2"  file_count:22640  delete_count:1  deleted_byte_count:30645  replica_placement:100  version:3  modified_at_second:1614877504 
          volume id:273  size:1074873432  collection:"collection2"  file_count:20511  delete_count:3  deleted_byte_count:46076  replica_placement:100  version:3  modified_at_second:1614884046 
          volume id:274  size:1075994128  collection:"collection2"  file_count:20997  replica_placement:100  version:3  modified_at_second:1614884113 
          volume id:276  size:1076899888  collection:"collection2"  file_count:20190  delete_count:1  deleted_byte_count:8798  replica_placement:100  version:3  modified_at_second:1614884003 
          volume id:277  size:1074956160  collection:"collection2"  file_count:19260  delete_count:2  deleted_byte_count:172356  replica_placement:100  version:3  modified_at_second:1614889988 
          volume id:279  size:1077325096  collection:"collection2"  file_count:19671  delete_count:6  deleted_byte_count:379116  replica_placement:100  version:3  modified_at_second:1614890230 
          volume id:282  size:1075232240  collection:"collection2"  file_count:22659  delete_count:4  deleted_byte_count:67915  replica_placement:100  version:3  modified_at_second:1614897304 
          volume id:284  size:1074533384  collection:"collection2"  file_count:22196  delete_count:4  deleted_byte_count:154683  replica_placement:100  version:3  modified_at_second:1614897231 
          volume id:285  size:1082128576  collection:"collection2"  file_count:21804  delete_count:1  deleted_byte_count:1064  replica_placement:100  version:3  modified_at_second:1614897165 
          volume id:286  size:1077464824  collection:"collection2"  file_count:23905  delete_count:6  deleted_byte_count:630577  replica_placement:100  version:3  modified_at_second:1614897401 
          volume id:287  size:1074590528  collection:"collection2"  file_count:28163  delete_count:5  deleted_byte_count:35727  replica_placement:100  version:3  modified_at_second:1614904874 
          volume id:288  size:1075406800  collection:"collection2"  file_count:27243  delete_count:2  deleted_byte_count:51519  replica_placement:100  version:3  modified_at_second:1614904738 
          volume id:292  size:1092010744  collection:"collection2"  file_count:26781  delete_count:5  deleted_byte_count:508910  replica_placement:100  version:3  modified_at_second:1614912327 
          volume id:293  size:1075409776  collection:"collection2"  file_count:26063  delete_count:4  deleted_byte_count:183834  replica_placement:100  version:3  modified_at_second:1614912235 
          volume id:294  size:1075443992  collection:"collection2"  file_count:26076  delete_count:4  deleted_byte_count:194914  replica_placement:100  version:3  modified_at_second:1614912220 
          volume id:295  size:1074702376  collection:"collection2"  file_count:24488  delete_count:3  deleted_byte_count:48555  replica_placement:100  version:3  modified_at_second:1614911929 
          volume id:300  size:1076212424  collection:"collection2"  file_count:22892  delete_count:2  deleted_byte_count:61320  replica_placement:100  version:3  modified_at_second:1614918464 
          volume id:304  size:1081038888  collection:"collection2"  file_count:24505  delete_count:2  deleted_byte_count:124447  replica_placement:100  version:3  modified_at_second:1614925567 
          volume id:305  size:1074185552  collection:"collection2"  file_count:22074  delete_count:5  deleted_byte_count:20221  replica_placement:100  version:3  modified_at_second:1614925312 
          volume id:310  size:1074761520  collection:"collection2"  file_count:21441  delete_count:3  deleted_byte_count:13934  replica_placement:100  version:3  modified_at_second:1614931077 
          volume id:311  size:1088248208  collection:"collection2"  file_count:23553  delete_count:6  deleted_byte_count:191716  replica_placement:100  version:3  modified_at_second:1614931460 
          volume id:312  size:1075037808  collection:"collection2"  file_count:22524  replica_placement:100  version:3  modified_at_second:1614937832 
          volume id:313  size:1074876016  collection:"collection2"  file_count:22404  delete_count:4  deleted_byte_count:51728  replica_placement:100  version:3  modified_at_second:1614937755 
          volume id:314  size:1074670840  collection:"collection2"  file_count:20964  delete_count:4  deleted_byte_count:304291  replica_placement:100  version:3  modified_at_second:1614937441 
          volume id:315  size:1084153456  collection:"collection2"  file_count:23638  delete_count:2  deleted_byte_count:53956  replica_placement:100  version:3  modified_at_second:1614937884 
          volume id:316  size:1077720784  collection:"collection2"  file_count:22605  delete_count:1  deleted_byte_count:8503  replica_placement:100  version:3  modified_at_second:1614937838 
          volume id:317  size:1076215040  collection:"collection2"  file_count:23572  delete_count:2  deleted_byte_count:1441356  replica_placement:100  version:3  modified_at_second:1614943965 
          volume id:319  size:1073952744  collection:"collection2"  file_count:22286  delete_count:2  deleted_byte_count:43421  replica_placement:100  version:3  modified_at_second:1614943810 
          volume id:320  size:1082437736  collection:"collection2"  file_count:21544  delete_count:3  deleted_byte_count:16712  replica_placement:100  version:3  modified_at_second:1614943591 
          volume id:321  size:1081477960  collection:"collection2"  file_count:23531  delete_count:5  deleted_byte_count:262564  replica_placement:100  version:3  modified_at_second:1614943982 
          volume id:322  size:1078471600  collection:"collection2"  file_count:21905  delete_count:3  deleted_byte_count:145002  replica_placement:100  version:3  modified_at_second:1614950574 
          volume id:324  size:1075606712  collection:"collection2"  file_count:20799  delete_count:1  deleted_byte_count:251210  replica_placement:100  version:3  modified_at_second:1614950310 
          volume id:326  size:1076059936  collection:"collection2"  file_count:22564  delete_count:2  deleted_byte_count:192886  replica_placement:100  version:3  modified_at_second:1614950619 
          volume id:327  size:1076121224  collection:"collection2"  file_count:22007  delete_count:3  deleted_byte_count:60358  replica_placement:100  version:3  modified_at_second:1614956487 
          volume id:328  size:1074767928  collection:"collection2"  file_count:21720  delete_count:3  deleted_byte_count:56429  replica_placement:100  version:3  modified_at_second:1614956362 
          volume id:329  size:1076691776  collection:"collection2"  file_count:22411  delete_count:5  deleted_byte_count:214092  replica_placement:100  version:3  modified_at_second:1614956485 
          volume id:331  size:1074957192  collection:"collection2"  file_count:21230  delete_count:4  deleted_byte_count:62145  replica_placement:100  version:3  modified_at_second:1614956259 
          volume id:333  size:1074270192  collection:"collection2"  file_count:21271  delete_count:2  deleted_byte_count:168122  replica_placement:100  version:3  modified_at_second:1614962697 
          volume id:335  size:1076235176  collection:"collection2"  file_count:22391  delete_count:3  deleted_byte_count:8838  replica_placement:100  version:3  modified_at_second:1614962970 
          volume id:336  size:1087853032  collection:"collection2"  file_count:22801  delete_count:2  deleted_byte_count:26394  replica_placement:100  version:3  modified_at_second:1614963003 
          volume id:338  size:1076118360  collection:"collection2"  file_count:21680  replica_placement:100  version:3  modified_at_second:1614969850 
          volume id:342  size:1080186296  collection:"collection2"  file_count:22405  delete_count:4  deleted_byte_count:64819  replica_placement:100  version:3  modified_at_second:1614977518 
          volume id:343  size:1075345184  collection:"collection2"  file_count:21095  delete_count:2  deleted_byte_count:20581  replica_placement:100  version:3  modified_at_second:1614977148 
          volume id:349  size:1075957824  collection:"collection2"  file_count:22395  delete_count:2  deleted_byte_count:61565  replica_placement:100  version:3  modified_at_second:1614984748 
          volume id:350  size:1074756688  collection:"collection2"  file_count:21990  delete_count:3  deleted_byte_count:233881  replica_placement:100  version:3  modified_at_second:1614984682 
          volume id:354  size:1085213992  collection:"collection2"  file_count:23150  delete_count:4  deleted_byte_count:82391  replica_placement:100  version:3  modified_at_second:1614992207 
          volume id:356  size:1083304992  collection:"collection2"  file_count:21552  delete_count:4  deleted_byte_count:14472  replica_placement:100  version:3  modified_at_second:1614992027 
          volume id:358  size:1085152312  collection:"collection2"  file_count:23756  delete_count:3  deleted_byte_count:44531  replica_placement:100  version:3  modified_at_second:1614998824 
          volume id:359  size:1074211240  collection:"collection2"  file_count:22437  delete_count:2  deleted_byte_count:187953  replica_placement:100  version:3  modified_at_second:1614998711 
          volume id:362  size:1074074120  collection:"collection2"  file_count:20595  delete_count:1  deleted_byte_count:112145  replica_placement:100  version:3  modified_at_second:1615004407 
          volume id:363  size:1078859496  collection:"collection2"  file_count:23177  delete_count:4  deleted_byte_count:9601  replica_placement:100  version:3  modified_at_second:1615004822 
          volume id:364  size:1081280816  collection:"collection2"  file_count:22686  delete_count:1  deleted_byte_count:84375  replica_placement:100  version:3  modified_at_second:1615004813 
          volume id:365  size:1075736632  collection:"collection2"  file_count:22193  delete_count:5  deleted_byte_count:259033  replica_placement:100  version:3  modified_at_second:1615004776 
          volume id:366  size:1075267272  collection:"collection2"  file_count:21856  delete_count:5  deleted_byte_count:138363  replica_placement:100  version:3  modified_at_second:1615004703 
          volume id:367  size:1076403648  collection:"collection2"  file_count:22995  delete_count:2  deleted_byte_count:36955  replica_placement:100  version:3  modified_at_second:1615010985 
          volume id:368  size:1074822016  collection:"collection2"  file_count:22252  delete_count:4  deleted_byte_count:3291946  replica_placement:100  version:3  modified_at_second:1615010878 
          volume id:369  size:1091472040  collection:"collection2"  file_count:23709  delete_count:4  deleted_byte_count:400876  replica_placement:100  version:3  modified_at_second:1615011019 
          volume id:370  size:1076040480  collection:"collection2"  file_count:22092  delete_count:2  deleted_byte_count:115388  replica_placement:100  version:3  modified_at_second:1615010877 
          volume id:371  size:1078806160  collection:"collection2"  file_count:22685  delete_count:2  deleted_byte_count:68905  replica_placement:100  version:3  modified_at_second:1615010994 
          volume id:372  size:1076193312  collection:"collection2"  file_count:22774  delete_count:1  deleted_byte_count:3495  replica_placement:100  version:3  modified_at_second:1615016911 
          volume id:374  size:1085011080  collection:"collection2"  file_count:23054  delete_count:2  deleted_byte_count:89034  replica_placement:100  version:3  modified_at_second:1615016917 
          volume id:375  size:1076140688  collection:"collection2"  file_count:21880  delete_count:2  deleted_byte_count:51103  replica_placement:100  version:3  modified_at_second:1615016787 
          volume id:378  size:959273824  collection:"collection2"  file_count:15031  replica_placement:100  version:3  modified_at_second:1615632323 
          volume id:379  size:1014108592  collection:"collection2"  file_count:15360  delete_count:8  deleted_byte_count:2524591  replica_placement:100  version:3  modified_at_second:1615632323 
          volume id:380  size:1010760464  collection:"collection2"  file_count:14920  delete_count:7  deleted_byte_count:134859  replica_placement:100  version:3  modified_at_second:1615632322 
        Disk hdd total size:274627838960 file_count:3607097 deleted_file:13594 deleted_bytes:4185524457 
      DataNode 192.168.1.1:8080 total size:274627838960 file_count:3607097 deleted_file:13594 deleted_bytes:4185524457 
    Rack DefaultRack total size:274627838960 file_count:3607097 deleted_file:13594 deleted_bytes:4185524457 
  DataCenter dc4 total size:274627838960 file_count:3607097 deleted_file:13594 deleted_bytes:4185524457 
  DataCenter dc5 hdd(volume:299/3000 active:299 free:2701 remote:0)
    Rack DefaultRack hdd(volume:299/3000 active:299 free:2701 remote:0)
      DataNode 192.168.1.5:8080 hdd(volume:299/3000 active:299 free:2701 remote:0)
        Disk hdd(volume:299/3000 active:299 free:2701 remote:0)
          volume id:5  size:293806008  file_count:1669  delete_count:2  deleted_byte_count:274334  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631342 
          volume id:11  size:1109224552  collection:"collection0"  file_count:44  replica_placement:100  version:3  modified_at_second:1615629606 
          volume id:18  size:1096678688  collection:"collection0"  file_count:88  delete_count:4  deleted_byte_count:8633  replica_placement:100  version:3  modified_at_second:1615631673 
          volume id:19  size:1096923792  collection:"collection0"  file_count:100  delete_count:10  deleted_byte_count:75779917  replica_placement:100  version:3  compact_revision:4  modified_at_second:1615630117 
          volume id:20  size:1074760432  collection:"collection0"  file_count:82  delete_count:5  deleted_byte_count:12156  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615629340 
          volume id:26  size:532657632  collection:"collection0"  file_count:100  delete_count:4  deleted_byte_count:9081  replica_placement:100  version:3  modified_at_second:1614170024 
          volume id:27  size:298886792  file_count:1608  replica_placement:100  version:3  modified_at_second:1615632482 
          volume id:28  size:308919192  file_count:1591  delete_count:1  deleted_byte_count:125280  replica_placement:100  version:3  modified_at_second:1615631762 
          volume id:29  size:281582688  file_count:1537  replica_placement:100  version:3  modified_at_second:1615629422 
          volume id:30  size:289466144  file_count:1566  delete_count:1  deleted_byte_count:124972  replica_placement:100  version:3  modified_at_second:1615632422 
          volume id:31  size:273363256  file_count:1498  replica_placement:100  version:3  modified_at_second:1615631642 
          volume id:32  size:281343360  file_count:1497  replica_placement:100  version:3  modified_at_second:1615632362 
          volume id:33  size:1130226344  collection:"collection1"  file_count:7322  delete_count:172  deleted_byte_count:45199399  replica_placement:100  version:3  modified_at_second:1615618789 
          volume id:34  size:1077111136  collection:"collection1"  file_count:9781  delete_count:110  deleted_byte_count:20894827  replica_placement:100  version:3  modified_at_second:1615619366 
          volume id:35  size:1075241744  collection:"collection1"  file_count:10523  delete_count:97  deleted_byte_count:46586217  replica_placement:100  version:3  modified_at_second:1615618790 
          volume id:36  size:1075118336  collection:"collection1"  file_count:10341  delete_count:118  deleted_byte_count:24753278  replica_placement:100  version:3  modified_at_second:1615606148 
          volume id:37  size:1075895576  collection:"collection1"  file_count:12013  delete_count:98  deleted_byte_count:50747932  replica_placement:100  version:3  modified_at_second:1615594776 
          volume id:38  size:1075545744  collection:"collection1"  file_count:13324  delete_count:100  deleted_byte_count:25223906  replica_placement:100  version:3  modified_at_second:1615569830 
          volume id:39  size:1076606536  collection:"collection1"  file_count:12612  delete_count:78  deleted_byte_count:17462730  replica_placement:100  version:3  modified_at_second:1615611959 
          volume id:40  size:1075358552  collection:"collection1"  file_count:12597  delete_count:62  deleted_byte_count:11657901  replica_placement:100  version:3  modified_at_second:1615612994 
          volume id:41  size:1076283592  collection:"collection1"  file_count:12088  delete_count:84  deleted_byte_count:19311237  replica_placement:100  version:3  modified_at_second:1615596736 
          volume id:42  size:1093948352  collection:"collection1"  file_count:7889  delete_count:47  deleted_byte_count:5697275  replica_placement:100  version:3  modified_at_second:1615548906 
          volume id:43  size:1116445864  collection:"collection1"  file_count:7355  delete_count:57  deleted_byte_count:9727158  replica_placement:100  version:3  modified_at_second:1615566167 
          volume id:44  size:1077582560  collection:"collection1"  file_count:7295  delete_count:50  deleted_byte_count:12618414  replica_placement:100  version:3  modified_at_second:1615566170 
          volume id:45  size:1075254640  collection:"collection1"  file_count:10772  delete_count:76  deleted_byte_count:22426345  replica_placement:100  version:3  modified_at_second:1615573498 
          volume id:46  size:1075286056  collection:"collection1"  file_count:9947  delete_count:309  deleted_byte_count:105601163  replica_placement:100  version:3  modified_at_second:1615569825 
          volume id:47  size:444599784  collection:"collection1"  file_count:709  delete_count:19  deleted_byte_count:11913451  replica_placement:100  version:3  modified_at_second:1615632397 
          volume id:48  size:1076778664  collection:"collection1"  file_count:9850  delete_count:77  deleted_byte_count:16641907  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630690 
          volume id:49  size:1078775288  collection:"collection1"  file_count:9631  delete_count:27  deleted_byte_count:5985628  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615575823 
          volume id:50  size:1076688288  collection:"collection1"  file_count:7921  delete_count:26  deleted_byte_count:5162032  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615610876 
          volume id:51  size:1076796120  collection:"collection1"  file_count:10550  delete_count:39  deleted_byte_count:12723654  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615547786 
          volume id:53  size:1063089216  collection:"collection1"  file_count:9832  delete_count:31  deleted_byte_count:9273066  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632156 
          volume id:54  size:1045022288  collection:"collection1"  file_count:9409  delete_count:29  deleted_byte_count:15102818  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630813 
          volume id:55  size:1012890016  collection:"collection1"  file_count:8651  delete_count:27  deleted_byte_count:9418841  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631453 
          volume id:56  size:1002412240  collection:"collection1"  file_count:8762  delete_count:40  deleted_byte_count:65885831  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615632014 
          volume id:57  size:839849792  collection:"collection1"  file_count:7514  delete_count:24  deleted_byte_count:6228543  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631775 
          volume id:58  size:908064192  collection:"collection1"  file_count:8128  delete_count:21  deleted_byte_count:6113731  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632343 
          volume id:59  size:988302272  collection:"collection1"  file_count:8098  delete_count:20  deleted_byte_count:3947615  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632238 
          volume id:60  size:1010702480  collection:"collection1"  file_count:8969  delete_count:79  deleted_byte_count:24782814  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615632439 
          volume id:61  size:975604544  collection:"collection1"  file_count:8683  delete_count:20  deleted_byte_count:10276072  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615631176 
          volume id:62  size:873845904  collection:"collection1"  file_count:7897  delete_count:23  deleted_byte_count:10920170  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631133 
          volume id:63  size:956941176  collection:"collection1"  file_count:8271  delete_count:32  deleted_byte_count:15876189  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632036 
          volume id:64  size:965638424  collection:"collection1"  file_count:8218  delete_count:27  deleted_byte_count:6922489  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631032 
          volume id:65  size:823283608  collection:"collection1"  file_count:7834  delete_count:29  deleted_byte_count:5950610  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632307 
          volume id:66  size:821343440  collection:"collection1"  file_count:7383  delete_count:29  deleted_byte_count:12010343  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631968 
          volume id:67  size:878713880  collection:"collection1"  file_count:7299  delete_count:117  deleted_byte_count:24857326  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632156 
          volume id:69  size:863913896  collection:"collection1"  file_count:7291  delete_count:100  deleted_byte_count:25335024  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630534 
          volume id:70  size:886695472  collection:"collection1"  file_count:7769  delete_count:164  deleted_byte_count:45162513  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632398 
          volume id:71  size:907608392  collection:"collection1"  file_count:7658  delete_count:122  deleted_byte_count:27622941  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632307 
          volume id:72  size:903990720  collection:"collection1"  file_count:6996  delete_count:240  deleted_byte_count:74147727  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630985 
          volume id:73  size:929047544  collection:"collection1"  file_count:7038  delete_count:227  deleted_byte_count:65336664  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630707 
          volume id:75  size:908045000  collection:"collection1"  file_count:6911  delete_count:268  deleted_byte_count:73934373  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632430 
          volume id:76  size:985296744  collection:"collection1"  file_count:6566  delete_count:61  deleted_byte_count:44464430  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632284 
          volume id:77  size:929398296  collection:"collection1"  file_count:7427  delete_count:238  deleted_byte_count:59581579  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632014 
          volume id:78  size:1075671512  collection:"collection1"  file_count:7540  delete_count:258  deleted_byte_count:71726846  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615582829 
          volume id:79  size:948225472  collection:"collection1"  file_count:6997  delete_count:227  deleted_byte_count:60625763  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631326 
          volume id:80  size:827912928  collection:"collection1"  file_count:6916  delete_count:15  deleted_byte_count:5611604  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615631159 
          volume id:81  size:880693168  collection:"collection1"  file_count:7481  delete_count:238  deleted_byte_count:80880878  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615631395 
          volume id:82  size:1041660512  collection:"collection1"  file_count:7043  delete_count:207  deleted_byte_count:52275724  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632430 
          volume id:83  size:936194288  collection:"collection1"  file_count:7593  delete_count:13  deleted_byte_count:4633917  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632029 
          volume id:84  size:871262320  collection:"collection1"  file_count:8190  delete_count:14  deleted_byte_count:3150948  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631161 
          volume id:86  size:1009434632  collection:"collection1"  file_count:8474  delete_count:236  deleted_byte_count:64543674  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630812 
          volume id:87  size:922274624  collection:"collection1"  file_count:12902  delete_count:13  deleted_byte_count:3412959  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615632438 
          volume id:88  size:1073767976  collection:"collection1"  file_count:14994  delete_count:207  deleted_byte_count:82380696  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615541383 
          volume id:89  size:1044421824  collection:"collection1"  file_count:14943  delete_count:243  deleted_byte_count:58543159  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615632208 
          volume id:90  size:891163760  collection:"collection1"  file_count:14608  delete_count:10  deleted_byte_count:2564369  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615629392 
          volume id:91  size:936573952  collection:"collection1"  file_count:14686  delete_count:11  deleted_byte_count:4717727  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631851 
          volume id:92  size:992439144  collection:"collection1"  file_count:7061  delete_count:195  deleted_byte_count:60649573  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630566 
          volume id:93  size:1079602592  collection:"collection1"  file_count:7878  delete_count:270  deleted_byte_count:74150048  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615556013 
          volume id:94  size:1030684704  collection:"collection1"  file_count:7660  delete_count:207  deleted_byte_count:70150733  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631616 
          volume id:95  size:990877824  collection:"collection1"  file_count:6620  delete_count:206  deleted_byte_count:60363604  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615631867 
          volume id:96  size:989294848  collection:"collection1"  file_count:7544  delete_count:229  deleted_byte_count:59931853  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630778 
          volume id:98  size:1077836472  collection:"collection1"  file_count:7605  delete_count:202  deleted_byte_count:73180379  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615523691 
          volume id:99  size:1071718496  collection:"collection1"  file_count:7470  delete_count:8  deleted_byte_count:9624950  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615631175 
          volume id:100  size:1083617472  collection:"collection1"  file_count:7018  delete_count:187  deleted_byte_count:61304236  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615505914 
          volume id:101  size:1077109408  collection:"collection1"  file_count:7706  delete_count:226  deleted_byte_count:77864780  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615630994 
          volume id:102  size:1074359920  collection:"collection1"  file_count:7338  delete_count:7  deleted_byte_count:6499151  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615626682 
          volume id:103  size:1075863904  collection:"collection1"  file_count:7184  delete_count:186  deleted_byte_count:58872238  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615628417 
          volume id:106  size:1075458680  collection:"collection1"  file_count:7182  delete_count:307  deleted_byte_count:69349053  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615598137 
          volume id:107  size:1073811776  collection:"collection1"  file_count:7436  delete_count:168  deleted_byte_count:57747428  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615293569 
          volume id:108  size:1074648024  collection:"collection1"  file_count:7472  delete_count:194  deleted_byte_count:70864699  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615593231 
          volume id:109  size:1075254560  collection:"collection1"  file_count:7556  delete_count:263  deleted_byte_count:55155265  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615502487 
          volume id:110  size:1076575744  collection:"collection1"  file_count:6996  delete_count:163  deleted_byte_count:52954032  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615590786 
          volume id:111  size:1073826176  collection:"collection1"  file_count:7355  delete_count:155  deleted_byte_count:50083578  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615593232 
          volume id:112  size:1076392512  collection:"collection1"  file_count:8291  delete_count:156  deleted_byte_count:74120183  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615569823 
          volume id:113  size:1076709184  collection:"collection1"  file_count:9355  delete_count:177  deleted_byte_count:59796765  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615569822 
          volume id:114  size:1074762792  collection:"collection1"  file_count:8802  delete_count:156  deleted_byte_count:38470055  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615591826 
          volume id:115  size:1076192296  collection:"collection1"  file_count:7690  delete_count:154  deleted_byte_count:32267193  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615285296 
          volume id:117  size:1073917192  collection:"collection1"  file_count:9520  delete_count:114  deleted_byte_count:21835126  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615573712 
          volume id:118  size:1074064344  collection:"collection1"  file_count:8738  delete_count:15  deleted_byte_count:3460697  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615516264 
          volume id:120  size:1076115928  collection:"collection1"  file_count:9639  delete_count:118  deleted_byte_count:33357871  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615482567 
          volume id:121  size:1078803320  collection:"collection1"  file_count:10113  delete_count:441  deleted_byte_count:94128627  replica_placement:100  version:3  modified_at_second:1615506626 
          volume id:122  size:1076235312  collection:"collection1"  file_count:9106  delete_count:252  deleted_byte_count:93041272  replica_placement:100  version:3  modified_at_second:1615585912 
          volume id:123  size:1080491112  collection:"collection1"  file_count:10623  delete_count:302  deleted_byte_count:83956998  replica_placement:100  version:3  modified_at_second:1615585916 
          volume id:124  size:1074519360  collection:"collection1"  file_count:9457  delete_count:286  deleted_byte_count:74752459  replica_placement:100  version:3  modified_at_second:1615585913 
          volume id:125  size:1088687040  collection:"collection1"  file_count:9518  delete_count:281  deleted_byte_count:76037905  replica_placement:100  version:3  modified_at_second:1615585913 
          volume id:126  size:1073867408  collection:"collection1"  file_count:9320  delete_count:278  deleted_byte_count:94547424  replica_placement:100  version:3  modified_at_second:1615585911 
          volume id:127  size:1074907336  collection:"collection1"  file_count:9900  delete_count:133  deleted_byte_count:48570820  replica_placement:100  version:3  modified_at_second:1615612990 
          volume id:128  size:1074874632  collection:"collection1"  file_count:9821  delete_count:148  deleted_byte_count:43633334  replica_placement:100  version:3  modified_at_second:1615602670 
          volume id:129  size:1074704328  collection:"collection1"  file_count:10012  delete_count:150  deleted_byte_count:64491721  replica_placement:100  version:3  modified_at_second:1615627566 
          volume id:130  size:1075000632  collection:"collection1"  file_count:10633  delete_count:161  deleted_byte_count:34768201  replica_placement:100  version:3  modified_at_second:1615582327 
          volume id:131  size:1075279584  collection:"collection1"  file_count:10075  delete_count:135  deleted_byte_count:29795712  replica_placement:100  version:3  modified_at_second:1615523898 
          volume id:132  size:1088539552  collection:"collection1"  file_count:11051  delete_count:71  deleted_byte_count:17178322  replica_placement:100  version:3  modified_at_second:1615619581 
          volume id:134  size:1074367304  collection:"collection1"  file_count:10662  delete_count:69  deleted_byte_count:25530139  replica_placement:100  version:3  modified_at_second:1615555873 
          volume id:135  size:1073906776  collection:"collection1"  file_count:10446  delete_count:71  deleted_byte_count:28599975  replica_placement:100  version:3  modified_at_second:1615569816 
          volume id:136  size:1074433552  collection:"collection1"  file_count:9593  delete_count:72  deleted_byte_count:26912512  replica_placement:100  version:3  modified_at_second:1615376036 
          volume id:137  size:1074309264  collection:"collection1"  file_count:9633  delete_count:50  deleted_byte_count:27487972  replica_placement:100  version:3  modified_at_second:1615572231 
          volume id:138  size:1074465744  collection:"collection1"  file_count:10120  delete_count:55  deleted_byte_count:15875438  replica_placement:100  version:3  modified_at_second:1615572231 
          volume id:140  size:1076203744  collection:"collection1"  file_count:11219  delete_count:57  deleted_byte_count:19864498  replica_placement:100  version:3  modified_at_second:1615571947 
          volume id:141  size:1074619488  collection:"collection1"  file_count:9840  delete_count:45  deleted_byte_count:40890181  replica_placement:100  version:3  modified_at_second:1615630994 
          volume id:142  size:1075733064  collection:"collection1"  file_count:9009  delete_count:48  deleted_byte_count:9912854  replica_placement:100  version:3  modified_at_second:1615598913 
          volume id:143  size:1075011280  collection:"collection1"  file_count:9608  delete_count:51  deleted_byte_count:37282460  replica_placement:100  version:3  modified_at_second:1615488584 
          volume id:144  size:1074549720  collection:"collection1"  file_count:8780  delete_count:50  deleted_byte_count:52475146  replica_placement:100  version:3  modified_at_second:1615573451 
          volume id:145  size:1074394928  collection:"collection1"  file_count:9255  delete_count:34  deleted_byte_count:38011392  replica_placement:100  version:3  modified_at_second:1615591825 
          volume id:146  size:1076337576  collection:"collection1"  file_count:10492  delete_count:50  deleted_byte_count:17071505  replica_placement:100  version:3  modified_at_second:1615632005 
          volume id:147  size:1077130576  collection:"collection1"  file_count:10451  delete_count:27  deleted_byte_count:8290907  replica_placement:100  version:3  modified_at_second:1615604115 
          volume id:148  size:1076066568  collection:"collection1"  file_count:9547  delete_count:33  deleted_byte_count:7034089  replica_placement:100  version:3  modified_at_second:1615586390 
          volume id:149  size:1074989016  collection:"collection1"  file_count:8352  delete_count:35  deleted_byte_count:7179742  replica_placement:100  version:3  modified_at_second:1615494494 
          volume id:150  size:1076290328  collection:"collection1"  file_count:9328  delete_count:33  deleted_byte_count:43417791  replica_placement:100  version:3  modified_at_second:1615611567 
          volume id:152  size:1075941400  collection:"collection1"  file_count:9951  delete_count:36  deleted_byte_count:25348335  replica_placement:100  version:3  modified_at_second:1615606614 
          volume id:153  size:1078539784  collection:"collection1"  file_count:10924  delete_count:34  deleted_byte_count:12603081  replica_placement:100  version:3  modified_at_second:1615606614 
          volume id:154  size:1081244696  collection:"collection1"  file_count:11002  delete_count:31  deleted_byte_count:8467560  replica_placement:100  version:3  modified_at_second:1615478469 
          volume id:155  size:1075140688  collection:"collection1"  file_count:10882  delete_count:32  deleted_byte_count:10076804  replica_placement:100  version:3  modified_at_second:1615606614 
          volume id:156  size:1074975832  collection:"collection1"  file_count:9535  delete_count:40  deleted_byte_count:11426621  replica_placement:100  version:3  modified_at_second:1615628341 
          volume id:157  size:1076758536  collection:"collection1"  file_count:10012  delete_count:19  deleted_byte_count:11688737  replica_placement:100  version:3  modified_at_second:1615597782 
          volume id:158  size:1087251976  collection:"collection1"  file_count:9972  delete_count:20  deleted_byte_count:10328429  replica_placement:100  version:3  modified_at_second:1615588879 
          volume id:159  size:1074132368  collection:"collection1"  file_count:9382  delete_count:27  deleted_byte_count:11474574  replica_placement:100  version:3  modified_at_second:1615593593 
          volume id:160  size:1075680952  collection:"collection1"  file_count:9772  delete_count:22  deleted_byte_count:4981968  replica_placement:100  version:3  modified_at_second:1615597780 
          volume id:162  size:1074286880  collection:"collection1"  file_count:11220  delete_count:17  deleted_byte_count:1815547  replica_placement:100  version:3  modified_at_second:1615478126 
          volume id:163  size:1074457192  collection:"collection1"  file_count:12524  delete_count:27  deleted_byte_count:6359619  replica_placement:100  version:3  modified_at_second:1615579313 
          volume id:164  size:1074261248  collection:"collection1"  file_count:11922  delete_count:25  deleted_byte_count:2923288  replica_placement:100  version:3  modified_at_second:1615620084 
          volume id:165  size:1073891016  collection:"collection1"  file_count:9152  delete_count:12  deleted_byte_count:19164659  replica_placement:100  version:3  modified_at_second:1615471907 
          volume id:166  size:1075637536  collection:"collection1"  file_count:14211  delete_count:24  deleted_byte_count:20415490  replica_placement:100  version:3  modified_at_second:1615491019 
          volume id:168  size:1074718808  collection:"collection1"  file_count:25702  delete_count:40  deleted_byte_count:4024775  replica_placement:100  version:3  modified_at_second:1615585664 
          volume id:169  size:1073863128  collection:"collection1"  file_count:25248  delete_count:43  deleted_byte_count:3013817  replica_placement:100  version:3  modified_at_second:1615569832 
          volume id:170  size:1075747096  collection:"collection1"  file_count:24596  delete_count:41  deleted_byte_count:3494711  replica_placement:100  version:3  modified_at_second:1615579204 
          volume id:171  size:1081881312  collection:"collection1"  file_count:24215  delete_count:36  deleted_byte_count:3191335  replica_placement:100  version:3  modified_at_second:1615596485 
          volume id:172  size:1074787312  collection:"collection1"  file_count:31236  delete_count:50  deleted_byte_count:3316482  replica_placement:100  version:3  modified_at_second:1615612385 
          volume id:173  size:1074154648  collection:"collection1"  file_count:30884  delete_count:34  deleted_byte_count:2430948  replica_placement:100  version:3  modified_at_second:1615591904 
          volume id:175  size:1077742504  collection:"collection1"  file_count:32353  delete_count:33  deleted_byte_count:1861403  replica_placement:100  version:3  modified_at_second:1615559515 
          volume id:176  size:1073854800  collection:"collection1"  file_count:30582  delete_count:34  deleted_byte_count:7701976  replica_placement:100  version:3  modified_at_second:1615626169 
          volume id:177  size:1074120120  collection:"collection1"  file_count:22293  delete_count:16  deleted_byte_count:3719562  replica_placement:100  version:3  modified_at_second:1615516891 
          volume id:178  size:1087560112  collection:"collection1"  file_count:23482  delete_count:22  deleted_byte_count:18810492  replica_placement:100  version:3  modified_at_second:1615541369 
          volume id:180  size:1078438536  collection:"collection1"  file_count:23614  delete_count:12  deleted_byte_count:4496474  replica_placement:100  version:3  modified_at_second:1614773242 
          volume id:181  size:1074571768  collection:"collection1"  file_count:22898  delete_count:19  deleted_byte_count:6628413  replica_placement:100  version:3  modified_at_second:1614745116 
          volume id:182  size:1076131280  collection:"collection1"  file_count:31987  delete_count:21  deleted_byte_count:1416142  replica_placement:100  version:3  modified_at_second:1615568922 
          volume id:183  size:1076361448  collection:"collection1"  file_count:31293  delete_count:16  deleted_byte_count:468841  replica_placement:100  version:3  modified_at_second:1615572982 
          volume id:184  size:1074594160  collection:"collection1"  file_count:31368  delete_count:22  deleted_byte_count:857453  replica_placement:100  version:3  modified_at_second:1615586578 
          volume id:185  size:1074099624  collection:"collection1"  file_count:30612  delete_count:17  deleted_byte_count:2610847  replica_placement:100  version:3  modified_at_second:1615506832 
          volume id:186  size:1074220864  collection:"collection1"  file_count:31450  delete_count:15  deleted_byte_count:391855  replica_placement:100  version:3  modified_at_second:1615614933 
          volume id:187  size:1074395944  collection:"collection1"  file_count:31853  delete_count:17  deleted_byte_count:454283  replica_placement:100  version:3  modified_at_second:1615590490 
          volume id:188  size:1074732792  collection:"collection1"  file_count:31867  delete_count:19  deleted_byte_count:393743  replica_placement:100  version:3  modified_at_second:1615487645 
          volume id:189  size:1074847896  collection:"collection1"  file_count:31450  delete_count:16  deleted_byte_count:1040552  replica_placement:100  version:3  modified_at_second:1615335661 
          volume id:190  size:1074008912  collection:"collection1"  file_count:31987  delete_count:11  deleted_byte_count:685125  replica_placement:100  version:3  modified_at_second:1615447161 
          volume id:191  size:1075493024  collection:"collection1"  file_count:31301  delete_count:19  deleted_byte_count:708401  replica_placement:100  version:3  modified_at_second:1615357456 
          volume id:192  size:1075857400  collection:"collection1"  file_count:31490  delete_count:25  deleted_byte_count:720617  replica_placement:100  version:3  modified_at_second:1615621632 
          volume id:193  size:1076616768  collection:"collection1"  file_count:31907  delete_count:16  deleted_byte_count:464900  replica_placement:100  version:3  modified_at_second:1615507875 
          volume id:194  size:1073985624  collection:"collection1"  file_count:31434  delete_count:18  deleted_byte_count:391432  replica_placement:100  version:3  modified_at_second:1615559499 
          volume id:195  size:1074158312  collection:"collection1"  file_count:31453  delete_count:15  deleted_byte_count:718266  replica_placement:100  version:3  modified_at_second:1615559331 
          volume id:196  size:1074594784  collection:"collection1"  file_count:31665  delete_count:18  deleted_byte_count:3468922  replica_placement:100  version:3  modified_at_second:1615501688 
          volume id:197  size:1075423296  collection:"collection1"  file_count:16473  delete_count:15  deleted_byte_count:12552442  replica_placement:100  version:3  modified_at_second:1615485253 
          volume id:198  size:1075104712  collection:"collection1"  file_count:16577  delete_count:18  deleted_byte_count:6583181  replica_placement:100  version:3  modified_at_second:1615623369 
          volume id:199  size:1078117688  collection:"collection1"  file_count:16497  delete_count:14  deleted_byte_count:1514286  replica_placement:100  version:3  modified_at_second:1615585984 
          volume id:200  size:1075630536  collection:"collection1"  file_count:16380  delete_count:18  deleted_byte_count:1103109  replica_placement:100  version:3  modified_at_second:1615485252 
          volume id:201  size:1091460440  collection:"collection1"  file_count:16684  delete_count:26  deleted_byte_count:5590335  replica_placement:100  version:3  modified_at_second:1615585987 
          volume id:202  size:1077533160  collection:"collection1"  file_count:2847  delete_count:67  deleted_byte_count:65172985  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615588497 
          volume id:203  size:1027316272  collection:"collection1"  file_count:3040  delete_count:11  deleted_byte_count:3993230  replica_placement:100  version:3  compact_revision:3  modified_at_second:1615631728 
          volume id:204  size:1079766872  collection:"collection1"  file_count:3233  delete_count:255  deleted_byte_count:104707641  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615565701 
          volume id:205  size:1078485304  collection:"collection1"  file_count:2869  delete_count:43  deleted_byte_count:18290259  replica_placement:100  version:3  compact_revision:2  modified_at_second:1615579314 
          volume id:206  size:1082045848  collection:"collection1"  file_count:2979  delete_count:225  deleted_byte_count:88220074  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630989 
          volume id:207  size:1081939960  collection:"collection1"  file_count:3010  delete_count:4  deleted_byte_count:692350  replica_placement:100  version:3  modified_at_second:1615269061 
          volume id:208  size:1077863624  collection:"collection1"  file_count:3147  delete_count:6  deleted_byte_count:858726  replica_placement:100  version:3  modified_at_second:1615495515 
          volume id:210  size:1094311304  collection:"collection1"  file_count:3468  delete_count:4  deleted_byte_count:466433  replica_placement:100  version:3  modified_at_second:1615495515 
          volume id:212  size:1078293448  collection:"collection1"  file_count:3106  delete_count:6  deleted_byte_count:2085755  replica_placement:100  version:3  modified_at_second:1615586387 
          volume id:213  size:1093588072  collection:"collection1"  file_count:3681  delete_count:12  deleted_byte_count:3138791  replica_placement:100  version:3  modified_at_second:1615586387 
          volume id:214  size:1074486992  collection:"collection1"  file_count:3217  delete_count:10  deleted_byte_count:6392871  replica_placement:100  version:3  modified_at_second:1615586383 
          volume id:215  size:1074798704  collection:"collection1"  file_count:2819  delete_count:31  deleted_byte_count:10873569  replica_placement:100  version:3  modified_at_second:1615586386 
          volume id:217  size:1075381872  collection:"collection1"  file_count:3331  delete_count:14  deleted_byte_count:2009141  replica_placement:100  version:3  modified_at_second:1615401638 
          volume id:218  size:1081263944  collection:"collection1"  file_count:3433  delete_count:14  deleted_byte_count:3454237  replica_placement:100  version:3  modified_at_second:1615603637 
          volume id:219  size:1092298816  collection:"collection1"  file_count:3193  delete_count:17  deleted_byte_count:2047576  replica_placement:100  version:3  modified_at_second:1615579316 
          volume id:220  size:1081928312  collection:"collection1"  file_count:3166  delete_count:13  deleted_byte_count:4127709  replica_placement:100  version:3  modified_at_second:1615579317 
          volume id:221  size:1106545456  collection:"collection1"  file_count:3153  delete_count:11  deleted_byte_count:1496835  replica_placement:100  version:3  modified_at_second:1615269138 
          volume id:222  size:1106623104  collection:"collection1"  file_count:3273  delete_count:11  deleted_byte_count:2114627  replica_placement:100  version:3  modified_at_second:1615586243 
          volume id:223  size:1075233064  collection:"collection1"  file_count:2966  delete_count:9  deleted_byte_count:744001  replica_placement:100  version:3  modified_at_second:1615586244 
          volume id:224  size:1093691520  collection:"collection1"  file_count:3463  delete_count:10  deleted_byte_count:1128328  replica_placement:100  version:3  modified_at_second:1615601870 
          volume id:225  size:1080698928  collection:"collection1"  file_count:3115  delete_count:7  deleted_byte_count:18170416  replica_placement:100  version:3  modified_at_second:1615434684 
          volume id:226  size:1103504768  collection:"collection1"  file_count:2965  delete_count:10  deleted_byte_count:2639254  replica_placement:100  version:3  modified_at_second:1615601867 
          volume id:228  size:1109784072  collection:"collection1"  file_count:2504  delete_count:24  deleted_byte_count:5458950  replica_placement:100  version:3  modified_at_second:1615610489 
          volume id:230  size:1080722984  collection:"collection1"  file_count:2898  delete_count:15  deleted_byte_count:3929261  replica_placement:100  version:3  modified_at_second:1615610537 
          volume id:232  size:1073901520  collection:"collection1"  file_count:3004  delete_count:54  deleted_byte_count:10273081  replica_placement:100  version:3  modified_at_second:1615611351 
          volume id:234  size:1073835280  collection:"collection1"  file_count:2965  delete_count:41  deleted_byte_count:4960354  replica_placement:100  version:3  modified_at_second:1615611351 
          volume id:235  size:1075586104  collection:"collection1"  file_count:2767  delete_count:33  deleted_byte_count:3216540  replica_placement:100  version:3  modified_at_second:1615611354 
          volume id:236  size:1089476136  collection:"collection1"  file_count:3231  delete_count:53  deleted_byte_count:11625921  replica_placement:100  version:3  modified_at_second:1615611351 
          volume id:237  size:375722792  collection:"collection1"  file_count:736  delete_count:16  deleted_byte_count:4464870  replica_placement:100  version:3  modified_at_second:1615631727 
          volume id:238  size:354320000  collection:"collection1"  file_count:701  delete_count:17  deleted_byte_count:5940420  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615632030 
          volume id:239  size:426569024  collection:"collection1"  file_count:693  delete_count:19  deleted_byte_count:13020783  replica_placement:100  version:3  compact_revision:1  modified_at_second:1615630841 
          volume id:240  size:424791528  collection:"collection1"  file_count:733  delete_count:13  deleted_byte_count:7515220  replica_placement:100  version:3  modified_at_second:1615631670 
          volume id:241  size:380217424  collection:"collection1"  file_count:633  delete_count:6  deleted_byte_count:1715768  replica_placement:100  version:3  modified_at_second:1615632006 
          volume id:242  size:1075383392  collection:"collection2"  file_count:10470  replica_placement:100  version:3  modified_at_second:1614852116 
          volume id:243  size:1088174704  collection:"collection2"  file_count:11109  delete_count:1  deleted_byte_count:938  replica_placement:100  version:3  modified_at_second:1614852203 
          volume id:244  size:1080295352  collection:"collection2"  file_count:10812  delete_count:1  deleted_byte_count:795  replica_placement:100  version:3  modified_at_second:1615628825 
          volume id:246  size:1075998672  collection:"collection2"  file_count:10365  delete_count:1  deleted_byte_count:13112  replica_placement:100  version:3  modified_at_second:1614852106 
          volume id:247  size:1075859808  collection:"collection2"  file_count:10443  delete_count:2  deleted_byte_count:564486  replica_placement:100  version:3  modified_at_second:1614856152 
          volume id:248  size:1084301208  collection:"collection2"  file_count:11217  delete_count:4  deleted_byte_count:746488  replica_placement:100  version:3  modified_at_second:1614856285 
          volume id:250  size:1080572168  collection:"collection2"  file_count:10220  replica_placement:100  version:3  modified_at_second:1614856129 
          volume id:252  size:1075065264  collection:"collection2"  file_count:14622  delete_count:2  deleted_byte_count:5228  replica_placement:100  version:3  modified_at_second:1614861200 
          volume id:253  size:1087328880  collection:"collection2"  file_count:14920  delete_count:3  deleted_byte_count:522994  replica_placement:100  version:3  modified_at_second:1614861258 
          volume id:254  size:1074830736  collection:"collection2"  file_count:14140  delete_count:2  deleted_byte_count:105892  replica_placement:100  version:3  modified_at_second:1614861115 
          volume id:255  size:1079581640  collection:"collection2"  file_count:14877  delete_count:3  deleted_byte_count:101223  replica_placement:100  version:3  modified_at_second:1614861233 
          volume id:256  size:1074283592  collection:"collection2"  file_count:14157  delete_count:1  deleted_byte_count:18156  replica_placement:100  version:3  modified_at_second:1614861100 
          volume id:257  size:1082621720  collection:"collection2"  file_count:18172  delete_count:2  deleted_byte_count:25125  replica_placement:100  version:3  modified_at_second:1614866402 
          volume id:258  size:1075527216  collection:"collection2"  file_count:18421  delete_count:4  deleted_byte_count:267833  replica_placement:100  version:3  modified_at_second:1614866420 
          volume id:259  size:1075507848  collection:"collection2"  file_count:18079  delete_count:2  deleted_byte_count:71992  replica_placement:100  version:3  modified_at_second:1614866381 
          volume id:260  size:1075105664  collection:"collection2"  file_count:17316  delete_count:4  deleted_byte_count:2015310  replica_placement:100  version:3  modified_at_second:1614866226 
          volume id:261  size:1076628592  collection:"collection2"  file_count:18355  delete_count:1  deleted_byte_count:1155  replica_placement:100  version:3  modified_at_second:1614866420 
          volume id:262  size:1078492584  collection:"collection2"  file_count:20390  delete_count:3  deleted_byte_count:287601  replica_placement:100  version:3  modified_at_second:1614871601 
          volume id:264  size:1081624192  collection:"collection2"  file_count:21151  replica_placement:100  version:3  modified_at_second:1614871629 
          volume id:265  size:1076401104  collection:"collection2"  file_count:19932  delete_count:2  deleted_byte_count:160823  replica_placement:100  version:3  modified_at_second:1614871543 
          volume id:266  size:1075617552  collection:"collection2"  file_count:20075  delete_count:1  deleted_byte_count:1039  replica_placement:100  version:3  modified_at_second:1614871526 
          volume id:267  size:1075699376  collection:"collection2"  file_count:21039  delete_count:3  deleted_byte_count:59956  replica_placement:100  version:3  modified_at_second:1614877294 
          volume id:270  size:1076876424  collection:"collection2"  file_count:22057  delete_count:1  deleted_byte_count:43916  replica_placement:100  version:3  modified_at_second:1614877469 
          volume id:271  size:1076992704  collection:"collection2"  file_count:22640  delete_count:1  deleted_byte_count:30645  replica_placement:100  version:3  modified_at_second:1614877504 
          volume id:272  size:1076145912  collection:"collection2"  file_count:21034  delete_count:2  deleted_byte_count:216564  replica_placement:100  version:3  modified_at_second:1614884139 
          volume id:273  size:1074873432  collection:"collection2"  file_count:20511  delete_count:3  deleted_byte_count:46076  replica_placement:100  version:3  modified_at_second:1614884046 
          volume id:274  size:1075994184  collection:"collection2"  file_count:20997  replica_placement:100  version:3  modified_at_second:1614884113 
          volume id:275  size:1078349024  collection:"collection2"  file_count:20808  delete_count:1  deleted_byte_count:1118  replica_placement:100  version:3  modified_at_second:1614884147 
          volume id:276  size:1076899880  collection:"collection2"  file_count:20190  delete_count:1  deleted_byte_count:8798  replica_placement:100  version:3  modified_at_second:1614884003 
          volume id:278  size:1078798632  collection:"collection2"  file_count:20597  delete_count:5  deleted_byte_count:400060  replica_placement:100  version:3  modified_at_second:1614890292 
          volume id:280  size:1077432160  collection:"collection2"  file_count:20286  delete_count:1  deleted_byte_count:879  replica_placement:100  version:3  modified_at_second:1614890262 
          volume id:281  size:1077581064  collection:"collection2"  file_count:20206  delete_count:3  deleted_byte_count:143964  replica_placement:100  version:3  modified_at_second:1614890237 
          volume id:282  size:1075232184  collection:"collection2"  file_count:22659  delete_count:4  deleted_byte_count:67915  replica_placement:100  version:3  modified_at_second:1614897304 
          volume id:283  size:1080178880  collection:"collection2"  file_count:19462  delete_count:7  deleted_byte_count:660407  replica_placement:100  version:3  modified_at_second:1614896623 
          volume id:286  size:1077464816  collection:"collection2"  file_count:23905  delete_count:6  deleted_byte_count:630577  replica_placement:100  version:3  modified_at_second:1614897401 
          volume id:287  size:1074590536  collection:"collection2"  file_count:28163  delete_count:5  deleted_byte_count:35727  replica_placement:100  version:3  modified_at_second:1614904875 
          volume id:288  size:1075406920  collection:"collection2"  file_count:27243  delete_count:2  deleted_byte_count:51519  replica_placement:100  version:3  modified_at_second:1614904738 
          volume id:289  size:1075284312  collection:"collection2"  file_count:29342  delete_count:5  deleted_byte_count:100454  replica_placement:100  version:3  modified_at_second:1614904977 
          volume id:290  size:1074723800  collection:"collection2"  file_count:28340  delete_count:4  deleted_byte_count:199064  replica_placement:100  version:3  modified_at_second:1614904924 
          volume id:292  size:1092010672  collection:"collection2"  file_count:26781  delete_count:5  deleted_byte_count:508910  replica_placement:100  version:3  modified_at_second:1614912325 
          volume id:295  size:1074702320  collection:"collection2"  file_count:24488  delete_count:3  deleted_byte_count:48555  replica_placement:100  version:3  modified_at_second:1614911929 
          volume id:296  size:1077824056  collection:"collection2"  file_count:26741  delete_count:4  deleted_byte_count:199906  replica_placement:100  version:3  modified_at_second:1614912301 
          volume id:297  size:1080229176  collection:"collection2"  file_count:23409  delete_count:5  deleted_byte_count:46268  replica_placement:100  version:3  modified_at_second:1614918481 
          volume id:298  size:1075410024  collection:"collection2"  file_count:23222  delete_count:2  deleted_byte_count:46110  replica_placement:100  version:3  modified_at_second:1614918474 
          volume id:302  size:1077559640  collection:"collection2"  file_count:23124  delete_count:7  deleted_byte_count:293111  replica_placement:100  version:3  modified_at_second:1614925500 
          volume id:304  size:1081038944  collection:"collection2"  file_count:24505  delete_count:2  deleted_byte_count:124447  replica_placement:100  version:3  modified_at_second:1614925569 
          volume id:305  size:1074185376  collection:"collection2"  file_count:22074  delete_count:5  deleted_byte_count:20221  replica_placement:100  version:3  modified_at_second:1614925312 
          volume id:306  size:1074763952  collection:"collection2"  file_count:22939  replica_placement:100  version:3  modified_at_second:1614925462 
          volume id:307  size:1076567912  collection:"collection2"  file_count:23377  delete_count:2  deleted_byte_count:25453  replica_placement:100  version:3  modified_at_second:1614931448 
          volume id:308  size:1074022336  collection:"collection2"  file_count:23086  delete_count:2  deleted_byte_count:2127  replica_placement:100  version:3  modified_at_second:1614931401 
          volume id:311  size:1088248344  collection:"collection2"  file_count:23553  delete_count:6  deleted_byte_count:191716  replica_placement:100  version:3  modified_at_second:1614931463 
          volume id:312  size:1075037528  collection:"collection2"  file_count:22524  replica_placement:100  version:3  modified_at_second:1614937831 
          volume id:313  size:1074875960  collection:"collection2"  file_count:22404  delete_count:4  deleted_byte_count:51728  replica_placement:100  version:3  modified_at_second:1614937755 
          volume id:316  size:1077720776  collection:"collection2"  file_count:22605  delete_count:1  deleted_byte_count:8503  replica_placement:100  version:3  modified_at_second:1614937838 
          volume id:318  size:1075965168  collection:"collection2"  file_count:22459  delete_count:2  deleted_byte_count:37778  replica_placement:100  version:3  modified_at_second:1614943862 
          volume id:322  size:1078471536  collection:"collection2"  file_count:21905  delete_count:3  deleted_byte_count:145002  replica_placement:100  version:3  modified_at_second:1614950572 
          volume id:323  size:1074608056  collection:"collection2"  file_count:21605  delete_count:4  deleted_byte_count:172090  replica_placement:100  version:3  modified_at_second:1614950526 
          volume id:325  size:1080701232  collection:"collection2"  file_count:21735  replica_placement:100  version:3  modified_at_second:1614950525 
          volume id:326  size:1076059920  collection:"collection2"  file_count:22564  delete_count:2  deleted_byte_count:192886  replica_placement:100  version:3  modified_at_second:1614950619 
          volume id:327  size:1076121304  collection:"collection2"  file_count:22007  delete_count:3  deleted_byte_count:60358  replica_placement:100  version:3  modified_at_second:1614956487 
          volume id:328  size:1074767816  collection:"collection2"  file_count:21720  delete_count:3  deleted_byte_count:56429  replica_placement:100  version:3  modified_at_second:1614956362 
          volume id:329  size:1076691960  collection:"collection2"  file_count:22411  delete_count:5  deleted_byte_count:214092  replica_placement:100  version:3  modified_at_second:1614956485 
          volume id:330  size:1080825760  collection:"collection2"  file_count:22464  delete_count:2  deleted_byte_count:15771  replica_placement:100  version:3  modified_at_second:1614956476 
          volume id:331  size:1074957256  collection:"collection2"  file_count:21230  delete_count:4  deleted_byte_count:62145  replica_placement:100  version:3  modified_at_second:1614956259 
          volume id:332  size:1075569928  collection:"collection2"  file_count:22097  delete_count:3  deleted_byte_count:98273  replica_placement:100  version:3  modified_at_second:1614962869 
          volume id:333  size:1074270160  collection:"collection2"  file_count:21271  delete_count:2  deleted_byte_count:168122  replica_placement:100  version:3  modified_at_second:1614962697 
          volume id:334  size:1075607880  collection:"collection2"  file_count:22546  delete_count:6  deleted_byte_count:101538  replica_placement:100  version:3  modified_at_second:1614962978 
          volume id:335  size:1076235136  collection:"collection2"  file_count:22391  delete_count:3  deleted_byte_count:8838  replica_placement:100  version:3  modified_at_second:1614962970 
          volume id:337  size:1075646896  collection:"collection2"  file_count:21934  delete_count:1  deleted_byte_count:3397  replica_placement:100  version:3  modified_at_second:1614969937 
          volume id:339  size:1078402392  collection:"collection2"  file_count:22309  replica_placement:100  version:3  modified_at_second:1614969995 
          volume id:340  size:1079462152  collection:"collection2"  file_count:22319  delete_count:4  deleted_byte_count:93620  replica_placement:100  version:3  modified_at_second:1614969977 
          volume id:341  size:1074448360  collection:"collection2"  file_count:21590  delete_count:5  deleted_byte_count:160085  replica_placement:100  version:3  modified_at_second:1614969858 
          volume id:343  size:1075345072  collection:"collection2"  file_count:21095  delete_count:2  deleted_byte_count:20581  replica_placement:100  version:3  modified_at_second:1614977148 
          volume id:346  size:1076464112  collection:"collection2"  file_count:22320  delete_count:4  deleted_byte_count:798258  replica_placement:100  version:3  modified_at_second:1614977511 
          volume id:347  size:1075145248  collection:"collection2"  file_count:22178  delete_count:1  deleted_byte_count:79392  replica_placement:100  version:3  modified_at_second:1614984727 
          volume id:348  size:1080623544  collection:"collection2"  file_count:21667  delete_count:1  deleted_byte_count:2443  replica_placement:100  version:3  modified_at_second:1614984604 
          volume id:349  size:1075957672  collection:"collection2"  file_count:22395  delete_count:2  deleted_byte_count:61565  replica_placement:100  version:3  modified_at_second:1614984748 
          volume id:351  size:1078795120  collection:"collection2"  file_count:23660  delete_count:3  deleted_byte_count:102141  replica_placement:100  version:3  modified_at_second:1614984816 
          volume id:352  size:1077145936  collection:"collection2"  file_count:22066  delete_count:1  deleted_byte_count:1018  replica_placement:100  version:3  modified_at_second:1614992130 
          volume id:353  size:1074897496  collection:"collection2"  file_count:21266  delete_count:2  deleted_byte_count:3105374  replica_placement:100  version:3  modified_at_second:1614991951 
          volume id:354  size:1085214104  collection:"collection2"  file_count:23150  delete_count:4  deleted_byte_count:82391  replica_placement:100  version:3  modified_at_second:1614992208 
          volume id:357  size:1074276152  collection:"collection2"  file_count:23137  delete_count:4  deleted_byte_count:188487  replica_placement:100  version:3  modified_at_second:1614998792 
          volume id:359  size:1074211296  collection:"collection2"  file_count:22437  delete_count:2  deleted_byte_count:187953  replica_placement:100  version:3  modified_at_second:1614998711 
          volume id:360  size:1075532512  collection:"collection2"  file_count:22574  delete_count:3  deleted_byte_count:1774776  replica_placement:100  version:3  modified_at_second:1614998770 
          volume id:361  size:1075362744  collection:"collection2"  file_count:22272  delete_count:1  deleted_byte_count:3497  replica_placement:100  version:3  modified_at_second:1614998668 
          volume id:362  size:1074074176  collection:"collection2"  file_count:20595  delete_count:1  deleted_byte_count:112145  replica_placement:100  version:3  modified_at_second:1615004407 
          volume id:363  size:1078859640  collection:"collection2"  file_count:23177  delete_count:4  deleted_byte_count:9601  replica_placement:100  version:3  modified_at_second:1615004823 
          volume id:364  size:1081280880  collection:"collection2"  file_count:22686  delete_count:1  deleted_byte_count:84375  replica_placement:100  version:3  modified_at_second:1615004813 
          volume id:365  size:1075736632  collection:"collection2"  file_count:22193  delete_count:5  deleted_byte_count:259033  replica_placement:100  version:3  modified_at_second:1615004776 
          volume id:366  size:1075267272  collection:"collection2"  file_count:21856  delete_count:5  deleted_byte_count:138363  replica_placement:100  version:3  modified_at_second:1615004703 
          volume id:367  size:1076403648  collection:"collection2"  file_count:22995  delete_count:2  deleted_byte_count:36955  replica_placement:100  version:3  modified_at_second:1615010985 
          volume id:368  size:1074821960  collection:"collection2"  file_count:22252  delete_count:4  deleted_byte_count:3291946  replica_placement:100  version:3  modified_at_second:1615010877 
          volume id:369  size:1091472040  collection:"collection2"  file_count:23709  delete_count:4  deleted_byte_count:400876  replica_placement:100  version:3  modified_at_second:1615011021 
          volume id:370  size:1076040544  collection:"collection2"  file_count:22092  delete_count:2  deleted_byte_count:115388  replica_placement:100  version:3  modified_at_second:1615010877 
          volume id:371  size:1078806216  collection:"collection2"  file_count:22685  delete_count:2  deleted_byte_count:68905  replica_placement:100  version:3  modified_at_second:1615010995 
          volume id:372  size:1076193344  collection:"collection2"  file_count:22774  delete_count:1  deleted_byte_count:3495  replica_placement:100  version:3  modified_at_second:1615016911 
          volume id:373  size:1080928088  collection:"collection2"  file_count:22617  delete_count:4  deleted_byte_count:91849  replica_placement:100  version:3  modified_at_second:1615016878 
          volume id:374  size:1085011176  collection:"collection2"  file_count:23054  delete_count:2  deleted_byte_count:89034  replica_placement:100  version:3  modified_at_second:1615016917 
          volume id:376  size:1074845832  collection:"collection2"  file_count:22908  delete_count:4  deleted_byte_count:432305  replica_placement:100  version:3  modified_at_second:1615016916 
          volume id:377  size:957434264  collection:"collection2"  file_count:14929  delete_count:1  deleted_byte_count:43099  replica_placement:100  version:3  modified_at_second:1615632323 
          volume id:379  size:1014108528  collection:"collection2"  file_count:15362  delete_count:6  deleted_byte_count:2481613  replica_placement:100  version:3  modified_at_second:1615632323 
        Disk hdd total size:306912958016 file_count:4201794 deleted_file:15268 deleted_bytes:4779359660 
      DataNode 192.168.1.5:8080 total size:306912958016 file_count:4201794 deleted_file:15268 deleted_bytes:4779359660 
    Rack DefaultRack total size:306912958016 file_count:4201794 deleted_file:15268 deleted_bytes:4779359660 
  DataCenter dc5 total size:306912958016 file_count:4201794 deleted_file:15268 deleted_bytes:4779359660 
total size:775256653592 file_count:10478712 deleted_file:33754 deleted_bytes:10839266043 
`
