import {
	Links,
	LiveReload,
	Meta,
	Outlet,
	Scripts,
	ScrollRestoration,
	useLoaderData,
} from '@remix-run/react'

import { HighlightInit } from '@highlight-run/remix/highlight-init'
import type { LinksFunction } from '@remix-run/node'
import { cssBundleHref } from '@remix-run/css-bundle'
import { json } from '@remix-run/node'
import { CONSTANTS } from '~/constants'

export { ErrorBoundary } from '~/components/error-boundary'

export const links: LinksFunction = () => [
	...(cssBundleHref ? [{ rel: 'stylesheet', href: cssBundleHref }] : []),
]

export async function loader() {
	return json({
		ENV: {
			HIGHLIGHT_PROJECT_ID: CONSTANTS.HIGHLIGHT_PROJECT_ID,
		},
	})
}

export default function App() {
	const { ENV } = useLoaderData()

	return (
		<html lang="en">
			<HighlightInit
				projectId={ENV.HIGHLIGHT_PROJECT_ID}
				tracingOrigins
				networkRecording={{ enabled: true, recordHeadersAndBody: true }}
			/>

			<head>
				<meta charSet="utf-8" />
				<meta
					name="viewport"
					content="width=device-width,initial-scale=1"
				/>
				<Meta />
				<Links />
			</head>
			<body>
				<Outlet />
				<ScrollRestoration />
				<script
					dangerouslySetInnerHTML={{
						__html: `window.ENV = ${JSON.stringify(ENV)}`,
					}}
				/>
				<Scripts />
				<LiveReload />
			</body>
		</html>
	)
}
