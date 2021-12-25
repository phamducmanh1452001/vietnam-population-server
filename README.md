# vietnam-population-server

Các api danh sách có chức năng search theo code và name với param key

Lấy danh sách tỉnh/thành phố: <br />
GET https://www.phorifai.xyz/api/provinces?page={int}&limit={int}&key={string} <br />

Lấy danh sách quận/huyện theo mã tỉnh/thành phố (province_code): <br />
GET https://www.phorifai.xyz/api/districts?province_code={string}&page={int}&limit={int}&key={string} <br />

Lấy danh sách phường/xã/thị trấn theo mã quận/huyện (district_code): <br />
GET https://www.phorifai.xyz/api/wards?district_code={string}&page={int}&limit={int}&key={string} <br />

Cán bộ đăng nhập (Mặc định password trùng code): <br />
POST https://www.phorifai.xyz/api/login <br />
Content-Type: application/x-www-form-urlencoded<br />
{ <br />
    "code": "YOUR_CODE", <br />
    "password": "YOUR_PASSWORD" <br />
} <br />

Cán bộ đăng xuất: <br />
POST https://www.phorifai.xyz/api/logout <br />
Authorization: Bearer "YOUR_TOKEN" <br />

Lấy danh sách cán bộ cấp dưới (dựa vào bearer token của cán bộ cấp trên): <br />
GET https://www.phorifai.xyz/api/lower-cadres&page={int}&limit={int}&key={string} <br />
Authorization: Bearer "YOUR_TOKEN" <br />

Lấy danh sách công dân (dựa vào bearer token của cán bộ quản lý): <br />
GET https://www.phorifai.xyz/api/citizens?page=2&limit=12 <br />
Authorization: Bearer "YOUR_TOKEN" <br />

Thay đổi quyền khai báo với cán bộ cấp dưới (0 hoặc 1): <br />
POST https://www.phorifai.xyz/api/change-cadre-permission <br />
Content-Type: application/x-www-form-urlencoded<br />
{ <br />
    "code": "CADRE_CODE", <br />
    "permission": 1 <br />
} <br />

Upload ảnh: (nên nén data ảnh <= 500KB)<br />
POST https://www.phorifai.xyz/api/upload <br />
Content-Type: multipart/form-data<br />
{<br />
    image: {Data}<br />
}<br />

Download ảnh:<br />
GET http://www.phorifai.xyz/api/images?name={string}
<br />
Ví dụ: http://www.phorifai.xyz/api/images?name=avatar1757368390.png
<br />
Thêm công dân:<br />
POST https://www.phorifai.xyz/api/add-citizen<br />
Authorization: Bearer "YOUR_TOKEN" <br />
{<br />
    "code": "808222771",<br />
    "first_name": "B",<br />
    "middle_name": "Van",<br />
    "last_name": "Pham",<br />
    "gender": "M",<br />
    "date_of_birth": "2001-01-01",<br />
    "date_of_joining": "2021-02-22",<br />
    "religion_id": 1,<br />
    "avatar": "",<br />
    "collaborator_name": "Quach Tinh",<br />
    "collaborator_phone": "0912345678"<br />
    "major": "Nghề nghiệp"<br />
    "temporary_address": "Địa chỉ tạm trú"<br />
}<br />

Xóa Công dân:<br />
DELETE https://www.phorifai.xyz/api/delete-citizen?code={string}<br />
Authorization: Bearer "YOUR_TOKEN" <br />
<br />
ID tôn giáo<br />
0: Không<br />
1: Phật giáo<br />
2: Công giáo<br />
3: Hồi giáo<br />
4: Tin Lành<br />
5: Cao Đài<br />
6: Tôn giáo dân gian<br />
7: Hòa Hảo<br />
8: Khác<br />
<br />
Lấy số liệu thống kê (code là mã khu vực tỉnh, huyện, xã, không truyền code nếu lấy cả nước)<br />
GET https://www.phorifai.xyz/api/age-chart?code=01<br />
GET https://www.phorifai.xyz/api/gender-chart?code=01<br />
GET https://www.phorifai.xyz/api/religion-chart?code=01<br />