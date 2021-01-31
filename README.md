# graytorrent
BitTorrent client implemented in golang

## In Progress...
- Manage multiple torrents
- Optimistic Unchoking
- Save the progress of managed torrents
- Fine-tune the adaptive download rate
- Limit number of connections

## Potential Goals
- Extensions (BEP6, DHT/PEX)
- Efficient rarest first requesting algorithm
- Use Google Protocol Buffers for network interactions

## Dependencies
- [bencode-go](https://github.com/jackpal/bencode-go)
- [pkg/errors](https://github.com/pkg/errors)
- [logrus](https://github.com/sirupsen/logrus)
- [pflag](https://github.com/spf13/pflag)
- [viper](https://github.com/spf13/viper)
- [testify](https://github.com/stretchr/testify)
- [logrus-prefixed-formatter](https://github.com/x-cray/logrus-prefixed-formatter)
