-- Con la base de datos "movies", se propone crear una tabla temporal llamada "TWD" 
-- y guardar en la misma los episodios de todas las temporadas de "The Walking Dead".


SELECT e.* FROM episodes e 
INNER JOIN seasons s ON e.season_id = s.id WHERE s.serie_id = (
SELECT id FROM series WHERE title = "The Walking Dead"
);

CREATE TEMPORARY TABLE `TWD`
SELECT e.* FROM episodes e 
INNER JOIN seasons s ON e.season_id = s.id WHERE s.serie_id = (
SELECT id FROM series WHERE title = "The Walking Dead"
);

-- Realizar una consulta a la tabla temporal para ver los episodios de la primera temporada.

SELECT * FROM TWD WHERE season_id IN (SELECT id FROM seasons WHERE title = "Primer Temporada");


-- En la base de datos "movies", seleccionar una tabla donde crear un índice 
-- y luego chequear la creación del mismo

ALTER TABLE series ADD INDEX `idx_series_title` (`title`);

SELECT * FROM series WHERE title = "The Walking Dead";

-- Analizar por qué crearía un índice en la tabla indicada y con qué criterio se elige/n el/los campos:

-- Se podria crear un indice si tenemos la necesidad de utilizar quizas un buscador mas eficiente. 
-- Los campos pueden ser quizas aquellos que identifiquen a los registros.