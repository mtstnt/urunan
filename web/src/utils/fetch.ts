

export async function fetchWithUser(url: string, info?: RequestInit) {
    const token = localStorage.getItem('token');
    if (token == null) {
        return null;
    }
    return fetch(url, {
        ...info,
        headers: {
            ...info?.headers,
            'Authorization': 'Bearer ' + token,
        },
    });
}