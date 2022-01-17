


-- name: HolaJony
-- var:id,var_type:int
-- var:height,var_type:float64
-- var:name
select * from TablaGeneral where id = ${{id}} and height = ${{height}} and name = "${{name}}"




-- name: HolaClaudio
-- var:id,var_type:int
-- var:lastName
select * from TablaScimuser where id = ${{id}} and  lastName = "${{lastName}}"