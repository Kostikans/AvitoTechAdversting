package advertisingRepository

const AddAdvertising = "INSERT INTO advertising VALUES(default,$1,$2,$3,$4,$5) RETURNING advertising_id"

const GetAdvertising = "SELECT name,cost,photos[1] from advertising where advertising_id=$1"

const listAdvertising = "SELECT name,photos[1],cost from advertising limit 10"
