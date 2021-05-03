import React from 'react'

export const Home: React.FC = () => {
  return (
    <main className="bg-grey-lighter min-h-screen flex flex-col">
      <div className="container max-w-sm mx-auto flex-1 flex flex-col items-center justify-center px-2">
        <div className="bg-white px-6 py-8 rounded shadow-md text-black w-full">
          <h1 className="mb-8 text-3xl text-center">Home</h1>
        </div>
      </div>
    </main>
  )
}
