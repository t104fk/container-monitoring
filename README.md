# container monitoring sample
Vagrantでアプリ用サーバーと監視用サーバーを作っている。  
アプリ用サーバーではGoのランタイムコンテナとcAdvisorのコンテナが動いている。  
cAdvisorのコンテナは監視用サーバーのinfluxdbにデータをぶちこんでいて、grafanaがそこからデータを引いてグラフ化している。  
詳しくは以下のURLを。  

### Special Thanks
- [cAdvisor, InfluxDB, GrafanaでDockerコンテナのリソース監視](http://qiita.com/atskimura/items/4c4aaaaa554e2814e938)
- [Exporting cAdvisor Stats to InfluxDB](https://github.com/google/cadvisor/blob/master/docs/influxdb.md)
