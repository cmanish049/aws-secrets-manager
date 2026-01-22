import { createBrowserRouter } from 'react-router-dom';
import Root from './routes/root';
import Dashboard, { loader as dashboardLoader } from './routes/dashboard';
import SecretView, { loader as secretLoader, action as secretAction } from './routes/secret';
import SecretNew, { action as newSecretAction } from './routes/secret-new';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    children: [
      {
        index: true,
        element: <Dashboard />,
        loader: dashboardLoader,
      },
      {
        path: 'secrets/new',
        element: <SecretNew />,
        action: newSecretAction,
      },
      {
        path: 'secrets/*',
        element: <SecretView />,
        loader: secretLoader,
        action: secretAction,
      },
    ],
  },
]);
