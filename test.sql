-- name: QueryTest1
-- var:id,var_type:int
-- var:height,var_type:float64
-- var:name,var_type:string,var_value:JAVA
select * from TablaGeneral
where id = ${{id}} and height = ${{height}}
and name = "${{name}}"

-- name: QueryTest2
-- var:id,var_type:int
-- var:lastName
select * from TablaScimuser
where id = ${{id}}
and  lastName = "${{lastName}}"

-- name: PrologQueryTest2
-- var:feature
getFeature("${{feature}}",X)