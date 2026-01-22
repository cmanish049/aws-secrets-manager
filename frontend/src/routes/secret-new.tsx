import { Form, redirect, useActionData, useNavigate } from 'react-router-dom';
import type { ActionFunctionArgs } from 'react-router-dom';
import { createSecret } from '../services/api';

export async function action({ request }: ActionFunctionArgs) {
  const formData = await request.formData();
  const name = formData.get('name') as string;
  const value = formData.get('value') as string;
  const description = formData.get('description') as string;

  try {
    await createSecret({ name, value, description });
    return redirect('/');
  } catch {
    return { error: 'Failed to create secret' };
  }
}

export default function SecretNew() {
  const actionData = useActionData() as { error?: string } | undefined;
  const navigate = useNavigate();

  return (
    <div className="bg-white shadow rounded-lg p-6">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold text-gray-900">Create New Secret</h1>
        <button
          onClick={() => navigate('/')}
          className="text-gray-600 hover:text-gray-900"
        >
          Cancel
        </button>
      </div>

      {actionData?.error && (
        <div className="mb-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
          {actionData.error}
        </div>
      )}

      <Form method="post" className="space-y-4">
        <div>
          <label htmlFor="name" className="block text-sm font-medium text-gray-700">
            Secret Name
          </label>
          <input
            type="text"
            id="name"
            name="name"
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="my-secret-name"
            required
          />
        </div>
        <div>
          <label htmlFor="value" className="block text-sm font-medium text-gray-700">
            Secret Value
          </label>
          <textarea
            id="value"
            name="value"
            rows={6}
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 font-mono text-sm"
            placeholder='{"key": "value"}'
            required
          />
        </div>
        <div>
          <label htmlFor="description" className="block text-sm font-medium text-gray-700">
            Description (optional)
          </label>
          <input
            type="text"
            id="description"
            name="description"
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="A brief description of this secret"
          />
        </div>
        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          Create Secret
        </button>
      </Form>
    </div>
  );
}
