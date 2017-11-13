package main

import (
	"testing"
)

func Test_hoge(t *testing.T) {
	script:=```
- method: http
  origin: resource/test.yaml 
  filter: 
	- name: genDate
	  key: Order.ServiceInfo.MVNOServiceno
	  format: pcdpEnteryyMMddss
	- name: genDate
	  key: Order.ServiceInfo.MVNOOrderNo
	  format: pcdpEnteryyMMddss

```

}
