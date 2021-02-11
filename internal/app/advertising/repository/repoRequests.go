package advertisingRepository

const AddAdvertising = "INSERT INTO advertising VALUES(default,$1,$2,$3,$4,$5) RETURNING advertising_id"

const GetAdvertising = "SELECT name,cost,photos[1] from advertising where advertising_id=$1"

const GetAdvertisingWithPhotos = "SELECT name,cost,photos[1],photos from advertising where advertising_id=$1"

const GetAdvertisingWithPhotosAndDescription = "SELECT name,cost,photos[1],photos,description  from advertising where advertising_id=$1"

const GetAdvertisingWithDescription = "SELECT name,cost,photos[1],description  from advertising where advertising_id=$1"

const CheckAdvertisingExist = "SELECT * from advertising where advertising_id=$1"

const GetPageCount = "SELECT count from advertising_count WHERE partition_id = (SELECT partition_id from advertising_count)"

const IncrementAdvertisingCount = "UPDATE advertising_count SET count = count + 1 WHERE partition_id = (SELECT partition_id from advertising_count)"
