-- Agregar una película a la tabla movies.
INSERT INTO movies (created_at, updated_at, title, rating, awards, release_date, length, genre_id) 
VALUES(now(), now(), "Spider Man 3", 9.8, 4,'2005-07-04 00:00:00', 300, 7);

-- Agregar un género a la tabla genres.
INSERT INTO genres (created_at, updated_at, name, ranking, active)
VALUES(now(), now(), "intense", 14, 1);

-- Asociar a la película del punto 1. con el género creado en el punto 2.
UPDATE movies
SET genre_id = 14
WHERE ID = 22;

-- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.
UPDATE actors
SET favorite_movie_id = 22
WHERE ID = 3;

-- Crear una tabla temporal copia de la tabla movies.

CREATE TEMPORARY TABLE temp_Movies
SELECT * FROM movies;

-- Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM temp_Movies 
WHERE awards < 5;

-- Obtener la lista de todos los géneros que tengan al menos una película.
SELECT * FROM genres 
INNER JOIN movies ON genre_id = movies.genre_id;

-- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
SELECT CONCAT(first_name, last_name) AS "name", movies.awards FROM actors
INNER JOIN movies ON actors.favorite_movie_id = movies.id
WHERE movies.awards > 3;

-- Crear un índice sobre el nombre en la tabla movies.
ALTER TABLE movies ADD INDEX `idx_movies_title` (`title`);

-- Chequee que el índice fue creado correctamente.
SELECT * FROM movies WHERE title = "Spider Man 3";

-- En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.
-- Seguramente si, en los campos que son propensos a ser buscados constantemente, por ejemplo 'titulo'.
-- Esto mejoraria la velocidad de consulta que posteriormente impactaria en la experiencia de usuario.
-- ¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
-- Quizas en la tabla episodes se podria implementar uno al buscar un episode por nombre. Justamente porque es probable que
-- sea un campo muy consultado y se podria optimizar la performance.