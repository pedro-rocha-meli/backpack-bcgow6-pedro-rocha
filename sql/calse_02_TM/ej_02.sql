/* En base al mismo, plantear las consultas SQL para resolver los siguientes requerimientos:

Listar los datos de los autores.
Listar nombre y edad de los estudiantes
¿Qué estudiantes pertenecen a la carrera informática?
¿Qué autores son de nacionalidad francesa o italiana?
¿Qué libros no son del área de internet?
Listar los libros de la editorial Salamandra.
Listar los datos de los estudiantes cuya edad es mayor al promedio.
Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
Listar los autores del libro "El Universo: Guía de viaje". (Se debe listar solamente los nombres).
¿Qué libros se prestaron al lector "Filippo Galli"?
Listar el nombre del estudiante de menor edad.
*/

-- a. Select de todos los campos de la tabla autor
-- b. Select de los campos nombre y edad de la tabla Estudiante
-- c. Select de todos los campos de la tabla Estudiante y Where que compare carrera con "informatica"
-- d. Select del campo nombre de la tabla Autor, Where que compare nacionalidad con "Francia" y OR que compare nacionalidad con "Italia"
-- e. Select del campo nombre de la tabla Libro, Where que compare area con un NOT LIKE "internet"
-- f. Select de todos los campos de la tabla Libro, Where que compare Editorial con un LIKE "Salamanca"
-- g. Select de todos los campos de la tabla Estudiantes, Where que compare que sea mayor la edad de un select anidado evaluando el promedio de la edad de los estudiantes
-- h. Select de nombre de la tabla Estudiantes, Where que compare apellido con un LIKE "G%"
-- i. Select de nombre de la tabla Autor, Join con la tabla Libro y where que compare nombre con un libro.Titulo LIKE "El Universo: Guía de viaje"
-- j. Falta la tabla Lector en el DER
-- k. Select de nombre de la tabla Estudiantes, where que compare edad con un select anidado evaluando con un MIN la edad de los estudiantes limit de 1