SELECT 
emp.id,
CONCAT(emp.name, ' ', emp.surname) AS full_name,
emp.salary,
emp.phone_number,
emp.email,
emp.gender,
pm.id AS pos_id,
pm.name AS pos_name,
dm.id AS dep_id,
dm.name AS dep_name
FROM employees AS emp
INNER JOIN department_masters AS dm ON emp.department_id = dm.id
INNER JOIN position_masters AS pm ON emp.position_id = pm.id
WHERE emp.deleted_at IS NULL AND emp.id = '12'