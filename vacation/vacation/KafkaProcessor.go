package vacation

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"fmt"
    "encoding/json"
    "vacation/config"
)

var producer *kafka.Producer
var consumer *kafka.Consumer
var topic string

func InitProducer(platform string){
	var err error
    var options = config.Reader(platform)
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": options["bootstrap_servers"]})
	if err != nil {
		panic(err)
	}
	topic = options["destination"]
}

func InitConsumer(platform string){
	var err error
    var options = config.Reader(platform)
	consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": options["bootstrap_servers"],
		"group.id":          options["group_id"],
		"broker.address.family": "v4",
		"session.timeout.ms":    6000,
		"auto.offset.reset": "earliest",
	})
	topic = options["destination"]

	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	KafkaConsumer()
	
}

func KafkaProducer() (*kafka.Producer, string){
	
	return producer, topic
}

func KafkaConsumer(){
    
	
    consumer.SubscribeTopics([]string{topic}, nil)

	var dat map[string]interface{}
    for {
        msg, err := consumer.ReadMessage(-1)
        if err == nil {
			if err := json.Unmarshal(msg.Value, &dat); err != nil {
				panic(err)
			}
            if dat["eventType"] == "VacationRegistered"{
                wheneverVacationRegistered_Use(dat)
            }
            if dat["eventType"] == "VacationCancelled"{
                wheneverVacationCancelled_Add(dat)
            }
            if dat["eventType"] == "VacationRejected"{
                wheneverVacationRejected_Add(dat)
            }
            if dat["eventType"] == "EmployeeJoined"{
                wheneverEmployeeJoined_RegisterUser(dat)
            }
            if dat["eventType"] == "VacationDaysInsufficient"{
                wheneverVacationDaysInsufficient_Update(dat)
            }
			if dat["eventType"] == "VacationDaysIntialized"{
				whenVacationDaysIntialized_then_CREATE_1(dat)
			}
			if dat["eventType"] == "VacationDaysAdded"{
				whenVacationDaysAdded_then_UPDATE_1(dat)
			}
			if dat["eventType"] == "VacationDaysUsed"{
				whenVacationDaysUsed_then_UPDATE_2(dat)
			}



			
        } else {
            // The client will automatically try to recover from all errors.
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
        }
    }
}

func Streamhandler(message string){
	producer, topic := KafkaProducer()
	
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

}