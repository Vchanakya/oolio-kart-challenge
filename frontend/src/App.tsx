import { Container, Title } from '@mantine/core';
import ProductList from './components/ProductList';

export default function App() {
  return (
    <Container>
      <Title order={1}>Products</Title>
      <ProductList />
    </Container>
  );
}
