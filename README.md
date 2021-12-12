# vietnam-population-server

Các api danh sách có chức năng search theo code và name với param key

Lấy danh sách tỉnh/thành phố: <br />
GET https://103.82.21.73/api/provinces?page={int}&limit={int}&key={string} <br />

Lấy danh sách quận/huyện theo mã tỉnh/thành phố (province_code): <br />
GET https://103.82.21.73/api/districts?province_code={string}&page={int}&limit={int}&key={string} <br />

Lấy danh sách phường/xã/thị trấn theo mã quận/huyện (district_code): <br />
GET https://103.82.21.73/api/wards?district_code={string}&page={int}&limit={int}&key={string} <br />

Cán bộ đăng nhập (Mặc định password trùng code): <br />
POST https://103.82.21.73/api/login <br />
Content-Type: application/x-www-form-urlencoded<br />
{ <br />
    "code": "YOUR_CODE", <br />
    "password": "YOUR_PASSWORD" <br />
} <br />

Cán bộ đăng xuất: <br />
POST https://103.82.21.73/api/logout <br />
Authorization: Bearer "YOUR_TOKEN" <br />

Lấy danh sách cán bộ cấp dưới (dựa vào bearer token của cán bộ cấp trên): <br />
GET https://103.82.21.73/api/lower-cadres&page={int}&limit={int}&key={string} <br />
Authorization: Bearer "YOUR_TOKEN" <br />

