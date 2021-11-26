# vietnam-population-server

Lấy danh sách tỉnh/thành phố:
GET https://intense-falls-59574.herokuapp.com/api/provinces

Lấy danh sách quận/huyện theo mã tỉnh/thành phố (province_code):
GET https://intense-falls-59574.herokuapp.com/api/districts?province_code={string}

Lấy danh sách phường/xã/thị trấn theo mã quận/huyện (district_code):
GET https://intense-falls-59574.herokuapp.com/api/wards?district_code={string}

Cán bộ đăng nhập (Mặc định password trùng code):
POST https://intense-falls-59574.herokuapp.com/api/login
Content-Type: application/x-www-form-urlencoded
{
    "code": "YOUR_CODE"
    "password": "YOUR_PASSWORD"    
}

Lấy danh sách cán bộ cấp dưới (dựa vào bearer token của cán bộ cấp trên):
GET https://intense-falls-59574.herokuapp.com/api/lower-cadres
Authorization: Bearer "YOUR_TOKEN"

