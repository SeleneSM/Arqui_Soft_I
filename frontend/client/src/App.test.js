import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import App from './App';

test('renders login component', () => {
  render(
    <BrowserRouter>
      <App />
    </BrowserRouter>
  );
  const userLabel = screen.getByLabelText(/user/i);
  expect(userLabel).toBeInTheDocument();
});
