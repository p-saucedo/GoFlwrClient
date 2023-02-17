# GoFlwrClient

## Project still under construction

[Flwr](https://flower.dev/) is one of the most widely used federated learning frameworks today, both in academic and productive environments. Flwr offers an interface on top of gRPC that allows cross-platform and multi-language clients and servers. The goal of this project is to develop a client written in Golang, which is perfectly compatible with the other available implementations, mainly in Python or NodeJS. 

The low level character of Flwr allows to use different ML libraries to train the same federated model, as long as a compatible data format is complied with in all nodes. This is why, for the PoC, [Gorgonia](https://gorgonia.org/getting-started/) has been used as the model development library in Golang, which allows the creation of both classical ML models and complex neural networks.

Any contribution will be welcome.

##### Resources used:
- https://github.com/Orange-OpenSource/flower-nodejs-client
- https://hackernoon.com/neural-networks-with-gorgonia-ag1a3r5a
