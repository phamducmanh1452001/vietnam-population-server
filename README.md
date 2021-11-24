# vietnam-population-server

Lấy danh sách tỉnh/thành phố
GET https://intense-falls-59574.herokuapp.com/provinces

Lấy danh sách quận/huyện theo mã tỉnh/thành phố (province_code)
GET https://intense-falls-59574.herokuapp.com/districts?province_code={string}

Lấy danh sách phường/xã/thị trấn theo mã quận/huyện (district_code)
GET https://intense-falls-59574.herokuapp.com/wards?district_code={string}
