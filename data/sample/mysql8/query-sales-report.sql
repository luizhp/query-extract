SELECT 
    orders.order_id AS order_id,
    sellers.first_name AS seller_first_name,
    sellers.last_name AS seller_last_name,
    products.product_name AS product_name,
    suppliers.company_name AS supplier_name,
    clients.first_name AS client_first_name,
    clients.last_name AS client_last_name,
    orders.order_date AS order_date,
    order_lines.quantity AS quantity,
    order_lines.unit_price AS unit_price,
    order_lines.discount AS discount,
    (order_lines.quantity * order_lines.unit_price * (1 - order_lines.discount)) AS line_total
FROM 
    order_lines
INNER JOIN 
    orders ON order_lines.order_id = orders.order_id
INNER JOIN 
    clients ON orders.client_id = clients.client_id
INNER JOIN 
    sellers ON orders.seller_id = sellers.seller_id
INNER JOIN 
    products ON order_lines.product_id = products.product_id
INNER JOIN 
    suppliers ON products.supplier_id = suppliers.supplier_id;