import { useState } from 'react';
import { useLoaderData, useNavigate, Form, useActionData, redirect } from 'react-router-dom';
import type { LoaderFunctionArgs, ActionFunctionArgs } from 'react-router-dom';
import { getSecret, updateSecret } from '../services/api';
import type { Secret } from '../services/api';

export async function loader({ params }: LoaderFunctionArgs): Promise<Secret> {
  const name = params['*']!;
  return await getSecret(name);
}

export async function action({ params, request }: ActionFunctionArgs) {
  const formData = await request.formData();
  const value = formData.get('value') as string;
  const name = params['*']!;

  try {
    await updateSecret(name, value);
    return redirect('/');
  } catch {
    return { error: 'Failed to update secret' };
  }
}

export default function SecretView() {
  const secret = useLoaderData() as Secret;
  const actionData = useActionData() as { error?: string } | undefined;
  const [isEditing, setIsEditing] = useState(false);
  const [value, setValue] = useState(secret.value || '');
  const navigate = useNavigate();

  const copyToClipboard = async () => {
    if (secret.value) {
      await navigator.clipboard.writeText(secret.value);
    }
  };

  return (
    <div className="bg-white shadow rounded-lg p-6">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold text-gray-900">{secret.name}</h1>
        <button
          onClick={() => navigate('/')}
          className="text-gray-600 hover:text-gray-900"
        >
          Back to list
        </button>
      </div>

      {actionData?.error && (
        <div className="mb-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
          {actionData.error}
        </div>
      )}

      {isEditing ? (
        <Form method="post" className="space-y-4">
          <div>
            <label htmlFor="value" className="block text-sm font-medium text-gray-700">
              Secret Value
            </label>
            <textarea
              id="value"
              name="value"
              rows={6}
              value={value}
              onChange={(e) => setValue(e.target.value)}
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 font-mono text-sm"
              required
            />
          </div>
          <div className="flex gap-4">
            <button
              type="submit"
              className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700"
            >
              Save Changes
            </button>
            <button
              type="button"
              onClick={() => {
                setIsEditing(false);
                setValue(secret.value || '');
              }}
              className="bg-gray-200 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-300"
            >
              Cancel
            </button>
          </div>
        </Form>
      ) : (
        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Secret Value
            </label>
            <div className="relative">
              <pre className="bg-gray-100 p-4 rounded-md overflow-x-auto font-mono text-sm whitespace-pre-wrap break-all">{secret.value}</pre>
              <button
                onClick={copyToClipboard}
                className="absolute top-2 right-2 bg-white px-3 py-1 rounded border border-gray-300 text-sm hover:bg-gray-50"
              >
                Copy
              </button>
            </div>
          </div>
          <button
            onClick={() => setIsEditing(true)}
            className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700"
          >
            Edit Secret
          </button>
        </div>
      )}
    </div>
  );
}
