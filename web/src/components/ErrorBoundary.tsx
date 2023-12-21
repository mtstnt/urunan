import React from 'react';

interface ErrorBoundaryState {
    hasError: boolean;
    error?: Error
}

export class CustomErrorBoundary extends React.Component<React.PropsWithChildren, ErrorBoundaryState> {

    constructor(props: React.PropsWithChildren) {
        super(props);
        this.state = { hasError: false };
    }

    componentDidCatch(error: Error, errorInfo: React.ErrorInfo): void {
        console.log("React Error: " + error, errorInfo);
    }

    static getDerivedStateFromError(error: Error) {
        return {
            hasError: true,
            error: error
        };
    }

    render(): React.ReactNode {
        if (this.state.hasError) {
            return <h1>Error!</h1>;
        } else {
            return this.props.children;
        }
    }

}