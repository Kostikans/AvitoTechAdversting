package advertisingRepository

const AddAdvertising = "INSERT INTO advertising VALUES(default,$1,$2,$3,$4) RETURNING advertising_id"

const GetAdvertising = "SELECT name,cost,photos[0] from advertising where advertising_id=$1"

const listAdvertising = "SELECT name.photos[0],cost from advertising limit 10"
