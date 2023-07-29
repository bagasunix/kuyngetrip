package village

const (
	QUERY_GET_COORDINATE = `SELECT * FROM (SELECT v.id as village_id, v.name as village_name, sub_district_id, s.name as sub_district_name, s.latitude as sub_district_latitude, s.longitude as sub_district_longitude, city_id, c.name as city_name, c.latitude as city_latitude, c.longitude as city_longitude, province_id, p.name as province_name, p.latitude as province_latitude, p.longitude as province_longitude, country_id, co.name as country_name,  co.latitude as country_latitude, co.longitude as country_longitude, v.postal_codes as village_postal_code, calculate_distance(s.latitude, s.longitude, @latitude, @longitude, 'K') as distance FROM athena.villages v inner join athena.sub_districts s on v.sub_district_id = s.id inner join athena.cities c on s.city_id = c.id inner join athena.provinces p on c.province_id = p.id inner join athena.countries co on p.country_id = co.id where v.name like @keyWord or s.name like @keyWord or c.name like @keyWord or p.name like @keyWord or v.postal_codes  like @keyWord OR s.postal_codes like @keyWord) a where distance < @radius order by distance, village_name limit @limited`
	QUERY_GET_KEYWORD    = `SELECT v.id as village_id, v.name as village_name, sub_district_id, s.name as sub_district_name, s.latitude as sub_district_latitude, s.longitude as sub_district_longitude, city_id, c.name as city_name, c.latitude as city_latitude, c.longitude as city_longitude, province_id, p.name as province_name, p.latitude as province_latitude, p.longitude as province_longitude, country_id, co.name as country_name,  co.latitude as country_latitude, co.longitude as country_longitude, v.postal_codes as village_postal_code FROM athena.villages v inner join athena.sub_districts s on v.sub_district_id = s.id inner join athena.cities c on s.city_id = c.id inner join athena.provinces p on c.province_id = p.id inner join athena.countries co on p.country_id = co.id where v.name like @keyWord or s.name like @keyWord or c.name like @keyWord or p.name like @keyWord or v.postal_codes  like @keyWord OR s.postal_codes like @keyWord order by v.name limit @limited`
)
