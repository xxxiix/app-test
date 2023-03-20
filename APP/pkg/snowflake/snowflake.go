package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
)

var node *sf.Node

func Init() (err error) {
	startTime := viper.GetString("app.start_time")
	machineID := viper.GetInt64("machine_id")
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}

// func main() {
// 	if err := Init("2020-07-01", 1); err != nil {
// 		fmt.Printf("init failed, err:%v\n", err)
// 		return
// 	}
// 	id := GenID()
// 	fmt.Println(id)
// }
