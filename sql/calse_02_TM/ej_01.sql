-- Adding Values to Departamento
INSERT INTO Departamento VALUES ("D-000-1", "Software", "Los Tigres");
INSERT INTO Departamento VALUES ("D-000-2", "Sistemas", "Guadalupe");
INSERT INTO Departamento VALUES ("D-000-3", "Contabilidad", "La Roca");
INSERT INTO Departamento VALUES ("D-000-4", "Ventas", "La Plata");


-- Adding Values to Empleado
INSERT INTO Empleado VALUES ("E-0001", "Cesar", "Cabarez", "Vendedor", "2018-05-12", 80000, 15000, "D-000-1");
INSERT INTO Empleado VALUES ("E-0002", "Marcos", "Brea", "Desarrollador", "2011-05-12", 80000, 15000, "D-000-2");
INSERT INTO Empleado VALUES ("E-0003", "Luana", "Fleming", "Analista", "2010-05-12", 80000, 15000, "D-000-3");
INSERT INTO Empleado VALUES ("E-0004", "Maria", "Tobarez", "Presidente", "2000-05-12", 80000, 15000, "D-000-4");
INSERT INTO Empleado VALUES ("E-0005", "Constanza", "Secorla", "Director", "2004-05-12", 80000, 15000, "D-000-2");
INSERT INTO Empleado VALUES ("E-0005", "Mito", "Barchuk", "Director", "2004-05-12", 80000, 15000, "D-000-4");



-- Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.

SELECT nombre, puesto, dep.localidad FROM Empleado emp 
INNER JOIN Departamento dep 
ON dep.depto_nro = emp.Departamento_depto_nro
WHERE emp.puesto = "Vendedor";

-- Visualizar los departamentos con más de cinco empleados.

SELECT * FROM Empleado emp 
INNER JOIN Departamento dep
ON emp.Departamento_depto_nro = dep.depto_nro
GROUP BY dep.nombre_depto HAVING COUNT(*) > 5;

-- Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.

SELECT nombre, salario, dep.nombre_depto 
FROM Empleado emp 
INNER JOIN Departamento dep 
ON emp.Departamento_depto_nro = dep.depto_nro
WHERE dep.nombre_depto = (SELECT puesto FROM Empleado emp WHERE emp.nombre = "Mito" AND emp.apellido = "Barchuk");

-- Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
SELECT * FROM Empleado emp
INNER JOIN Departamento dep
ON dep.depto_nro = emp.Departamento_depto_nro
WHERE dep.nombre_depto = "Contabilidad";

-- Mostrar el nombre del empleado que tiene el salario más bajo.

SELECT nombre FROM Empleado WHERE salario = (SELECT MIN(salario) FROM Empleado);