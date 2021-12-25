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
{
    image: {Data}
}

Download ảnh:
GET http://www.phorifai.xyz/api/images?name={string}

Ví dụ: http://www.phorifai.xyz/api/images?name=avatar1757368390.png

Thêm công dân:
POST https://www.phorifai.xyz/api/add-citizen
Authorization: Bearer "YOUR_TOKEN" <br />
{
    "code": "808222771",
    "first_name": "B",
    "middle_name": "Van",
    "last_name": "Pham",
    "gender": "M",
    "date_of_birth": "2001-01-01",
    "date_of_joining": "2021-02-22",
    "religion_id": 1,
    "avatar": "",
    "collaborator_name": "Quach Tinh",
    "collaborator_phone": "0912345678"
    "major": "Nghề nghiệp"
    "temporary_address": "Địa chỉ tạm trú"
}

ID tôn giáo
0: Không
1: Phật giáo
2: Công giáo
3: Hồi giáo
4: Tin Lành
5: Cao Đài
6: Tôn giáo dân gian
7: Hòa Hảo
8: Khác
<br />
Lấy số liệu thống kê (code là mã khu vực tỉnh, huyện, xã, không truyền code nếu lấy cả nước)<br />
GET https://www.phorifai.xyz/api/age-chart?code=01<br />
GET https://www.phorifai.xyz/api/gender-chart?code=01<br />
GET https://www.phorifai.xyz/api/religion-chart?code=01<br />