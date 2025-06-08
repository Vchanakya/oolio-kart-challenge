import { useEffect, useState } from 'react';
import { List, Loader } from '@mantine/core';

interface Product {
  id: string;
  name: string;
}

export default function ProductList() {
  const [products, setProducts] = useState<Product[] | null>(null);

  useEffect(() => {
    fetch('/product')
      .then((res) => res.json())
      .then(setProducts)
      .catch(() => setProducts([]));
  }, []);

  if (products === null) {
    return <Loader />;
  }

  return (
    <List>
      {products.map((p) => (
        <List.Item key={p.id}>{p.name}</List.Item>
      ))}
    </List>
  );
}
