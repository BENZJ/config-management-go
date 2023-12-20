# 问题清单
> unsupported Scan, storing driver.Value type []uint8 into type *time.Time

https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time

使用db, err := sqlx.Connect("mysql", "myuser:mypass@tcp(127.0.0.1:3306)/mydb?parseTime=true")解决问题