SELECT 
    orders.order_id AS id_pedido,
    sellers.first_name AS nome_vendedor,
    sellers.last_name AS sobrenome_vendedor,
    products.product_name AS nome_produto,
    suppliers.company_name AS nome_fornecedor,
    clients.first_name AS nome_cliente,
    clients.last_name AS sobrenome_cliente,
    orders.order_date AS data_pedido,
    order_lines.quantity AS quantidade,
    order_lines.unit_price AS preco_unitario,
    order_lines.discount AS desconto,
    (order_lines.quantity * order_lines.unit_price * (1 - order_lines.discount)) AS total_linha
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