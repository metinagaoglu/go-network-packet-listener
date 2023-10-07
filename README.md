### Simple Network Sniffer - (experimental)

This project is a simple network sniffer implemented in Golang. It captures packets from a specified network interface and streams them to a Kafka topic. The packets can then be consumed and processed in real-time.
Prerequisites

    Golang
    gopacket library
    sarama library for Kafka
    A running Kafka cluster


## Prerequisites

Before you get started, make sure you have the following prerequisites installed on your system:

- Go: You should have Go installed on your machine.

- Apache Kafka: You should have an Apache Kafka cluster up and running.


### Note

Please be aware that capturing and processing network packets can be a sensitive operation and should be done with caution. Make sure to have the appropriate permissions and understand the implications of capturing and processing network traffic before running this application.

## Resources

https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket

https://www.honeybadger.io/blog/golang-logging/

https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

https://medium.com/rungo/understanding-the-context-package-b2e407a9cdae
