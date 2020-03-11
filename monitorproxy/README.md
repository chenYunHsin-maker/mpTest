# monitorproxy

## Use kibana to query elasticsearch

1.Open web browser to kibana

2.Go to Dev Tools > Console

|API|Description|
|---|-----------|
| GET /_cat/indices | Get all elasticsearch indices |
| GET /event-vpn | Get event-vpn attributes |
| GET /event-vpn/_search?pretty=true&q=\*:\* | Default all event-vpn |
| GET /metric-reportvpntrafficinfo | Get metric-reportvpntrafficinfo attributes |
| GET /metric-reportvpntrafficinfo/_search?pretty=true&q=\*:\* | Default all metric-reportvpntrafficinfo |
| GET /metric-reportlinkquality | Get metric-reportlinkquality attributes |
| GET /metric-reportlinkquality/_search?pretty=true&q=\*:\* | Default all metric-reportlinkquality |
| GET /metric-reportliveinfo | GET metric-reportliveinfo attributes |
| GET /metric-reportliveinfo/_search?pretty=true&q=\*:\* | Default all metric-reportliveinfo |
| GET /metric-reporttrafficinfo | GET metric-reporttrafficinfo attributes |
| GET /metric-reporttrafficinfo/_search?pretty=true&q=\*:\* | Default all metric-reporttrafficinfo |
