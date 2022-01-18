
--name: QueryTest1
--var:id:number
--var:height:int
--var:name=hola como estas
select * from TablaGeneral
where id = ${{id}} and height = ${{height}}
and name = "${{name}}"

--name: QueryTest2
--var:id:float
--var:lastName:string
select * from TablaScimuser
where id = ${{id}}
and  lastName = "${{lastName}}"
