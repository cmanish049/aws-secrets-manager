import { useState } from 'react';
import { useLoaderData, Link, useNavigate } from 'react-router-dom';
import { listSecrets } from '../services/api';
import type { Secret } from '../services/api';

export async function loader(): Promise<Secret[]> {
  return await listSecrets();
}

export default function Dashboard() {
  const secrets = useLoaderData() as Secret[];
  const [search, setSearch] = useState('');
  const navigate = useNavigate();

  const filteredSecrets = secrets.filter((secret) =>
    secret.name.toLowerCase().includes(search.toLowerCase())
  );

  return (
    <div>
      <div className="mb-6">
        <input
          type="text"
          placeholder="Search secrets..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        />
      </div>

      {filteredSecrets.length === 0 ? (
        <div className="text-center py-12">
          <p className="text-gray-500">
            {search ? 'No secrets found matching your search.' : 'No secrets found.'}
          </p>
          <Link
            to="/secrets/new"
            className="mt-4 inline-block text-blue-600 hover:text-blue-700"
          >
            Create your first secret
          </Link>
        </div>
      ) : (
        <div className="bg-white shadow overflow-hidden rounded-md">
          <ul className="divide-y divide-gray-200">
            {filteredSecrets.map((secret) => (
              <li key={secret.name}>
                <button
                  onClick={() => navigate(`/secrets/${secret.name}`)}
                  className="w-full px-6 py-4 flex items-center justify-between hover:bg-gray-50 text-left"
                >
                  <div>
                    <p className="text-sm font-medium text-gray-900">{secret.name}</p>
                    {secret.description && (
                      <p className="text-sm text-gray-500">{secret.description}</p>
                    )}
                  </div>
                  <svg
                    className="h-5 w-5 text-gray-400"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M9 5l7 7-7 7"
                    />
                  </svg>
                </button>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}
