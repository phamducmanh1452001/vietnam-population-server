# vietnam-population-server

Lấy danh sách tỉnh/thành phố: <br />
GET https://evening-castle-31041.herokuapp.com/api/provinces <br />

Lấy danh sách quận/huyện theo mã tỉnh/thành phố (province_code): <br />
GET https://evening-castle-31041.herokuapp.com/api/districts?province_code={string} <br />

Lấy danh sách phường/xã/thị trấn theo mã quận/huyện (district_code): <br />
GET https://evening-castle-31041.herokuapp.com/api/wards?district_code={string} <br />

Cán bộ đăng nhập (Mặc định password trùng code): <br />
POST https://evening-castle-31041.herokuapp.com/api/login <br />
Content-Type: application/x-www-form-urlencoded<br />
{ <br />
    "code": "YOUR_CODE", <br />
    "password": "YOUR_PASSWORD" <br />
} <br />

Lấy danh sách cán bộ cấp dưới (dựa vào bearer token của cán bộ cấp trên): <br />
GET https://evening-castle-31041.herokuapp.com/api/lower-cadres <br />
Authorization: Bearer "YOUR_TOKEN" <br />

