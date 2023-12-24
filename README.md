# go-advantage-air

Go client for the Advantage Air system

### Usage

```bash
go get -u github.com/axatol/go-advantage-air
```

### Example

```go
package main

import advantageair "github.com/axatol/go-advantage-air"

func main() {
  ctx := context.Background()
  client := advantageair.NewClient("http://192.168.1.2:2025")

  data, err := client.GetSystemData(ctx)
  if err != nil {
    panic(err)
  }

  zone := data.GetAirconByID("ac1").GetZoneByName("z01")
  if zone.State == advantageair.AirconZoneStateOpen {
    change := advantageair.
      NewChange().
      SetAirconZoneState(advantageair.AirconZoneStateClose)

    if err := client.SetAircon(ctx, change); err != nil {
      panic(err)
    }
  }
}
```
