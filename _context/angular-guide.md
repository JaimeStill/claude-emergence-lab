# Angular v20 Modern Features & Best Practices Guide

## Why Angular for Greenfield Projects

Angular v20 represents a mature, production-ready framework with built-in solutions for common web application challenges. Unlike React's library approach requiring multiple third-party dependencies, Angular provides a complete development platform out of the box.

## Core Modern Features

### 1. Standalone Components (Default)

Angular v20 defaults to standalone components, eliminating the need for NgModules in most cases:

```typescript
@Component({
  selector: 'app-user-profile',
  imports: [CommonModule, ReactiveFormsModule],
  template: `...`
})
export class UserProfileComponent {}
```

### 2. Signals for Reactive State Management

Native reactivity without external libraries:

```typescript
@Component({
  template: `
    <p>Count: {{ count() }}</p>
    <p>Double: {{ doubleCount() }}</p>
  `
})
export class CounterComponent {
  count = signal(0);
  doubleCount = computed(() => this.count() * 2);
  
  increment() {
    this.count.update(v => v + 1);
  }
}
```

### 3. Control Flow Syntax

Modern template syntax replacing structural directives:

```typescript
@Component({
  template: `
    @if (user()) {
      <h1>Welcome, {{ user().name }}</h1>
    } @else {
      <p>Please log in</p>
    }
    
    @for (item of items(); track item.id) {
      <li>{{ item.name }}</li>
    } @empty {
      <p>No items found</p>
    }
    
    @switch (status()) {
      @case ('loading') { <app-spinner /> }
      @case ('error') { <app-error /> }
      @default { <app-content /> }
    }
  `
})
```

### 4. Zoneless Change Detection

Improved performance by eliminating Zone.js overhead:

```typescript
bootstrapApplication(AppComponent, {
  providers: [
    provideExperimentalZonelessChangeDetection()
  ]
});
```

### 5. Enhanced Template Expressions

JavaScript features now available in templates:

- Exponentiation operator (`**`)
- Template literals
- `in` operator for property checking
- `void` operator

## Best Practices

### State Management with Signals

```typescript
@Injectable({ providedIn: 'root' })
export class TodoService {
  private todos = signal<Todo[]>([]);
  
  // Expose readonly signal
  readonly todos$ = this.todos.asReadonly();
  
  // Computed values
  readonly completedCount = computed(() => 
    this.todos().filter(t => t.completed).length
  );
  
  addTodo(todo: Todo) {
    this.todos.update(todos => [...todos, todo]);
  }
}
```

### Component Architecture

```typescript
// Simplified naming convention in v20
// File: user.ts (not user.component.ts)
@Component({
  selector: 'app-user',
  imports: [CommonModule, UserAvatarComponent],
  template: `
    @let fullName = user().firstName + ' ' + user().lastName;
    
    <div class="user-card">
      <app-user-avatar [user]="user()" />
      <h2>{{ fullName }}</h2>
    </div>
  `
})
export class User {
  user = input.required<UserModel>();
}
```

### Lazy Loading with Standalone Components

```typescript
export const routes: Routes = [
  {
    path: 'products',
    loadComponent: () => import('./products/list').then(m => m.ProductList)
  },
  {
    path: 'products/:id',
    loadComponent: () => import('./products/detail').then(m => m.ProductDetail)
  }
];
```

### Signal-Based Forms

```typescript
@Component({
  template: `
    <form (submit)="onSubmit()">
      <input 
        [value]="formData().name" 
        (input)="updateName($event)"
      />
      
      @if (errors().name) {
        <span class="error">{{ errors().name }}</span>
      }
    </form>
  `
})
export class ContactForm {
  formData = signal({ name: '', email: '' });
  errors = computed(() => this.validateForm(this.formData()));
  
  updateName(event: Event) {
    const value = (event.target as HTMLInputElement).value;
    this.formData.update(data => ({ ...data, name: value }));
  }
}
```

### Performance Optimization

```typescript
// Use track for optimal rendering
@Component({
  template: `
    @for (product of products(); track product.id) {
      <app-product-card [product]="product" />
    }
  `
})
export class ProductList {
  products = signal<Product[]>([]);
}

// Defer loading heavy components
@Component({
  template: `
    @defer (on viewport) {
      <app-heavy-component />
    } @placeholder {
      <app-loading-skeleton />
    }
  `
})
export class DashboardComponent {}
```

### Testing with Signals

```typescript
describe('TodoService', () => {
  let service: TodoService;
  
  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TodoService);
  });
  
  it('should add todo', () => {
    const todo = { id: 1, text: 'Test', completed: false };
    
    service.addTodo(todo);
    
    // Synchronous assertion - no async needed!
    expect(service.todos$()).toContain(todo);
    expect(service.completedCount()).toBe(0);
  });
});
```

### Application Configuration & Providers

```typescript
// app.config.ts
import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { provideAnimations } from '@angular/platform-browser/animations';

import { routes } from './app.routes';
import { authInterceptor } from './core/interceptors/auth.interceptor';
import { errorInterceptor } from './core/interceptors/error.interceptor';

export const appConfig: ApplicationConfig = {
  providers: [
    // Use signals for change detection (zoneless)
    provideExperimentalZonelessChangeDetection(),
    // Or traditional zone-based detection
    // provideZoneChangeDetection({ eventCoalescing: true }),
    
    // Routing
    provideRouter(routes),
    
    // HTTP with interceptors
    provideHttpClient(
      withInterceptors([authInterceptor, errorInterceptor])
    ),
    
    // Animations
    provideAnimations(),
    
    // Custom providers
    { provide: API_BASE_URL, useValue: 'https://api.example.com' },
  ]
};
```

### Modern Routing Patterns

```typescript
// app.routes.ts
import { Routes } from '@angular/router';
import { authGuard } from './core/guards/auth.guard';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('./features/home/home').then(m => m.Home),
  },
  {
    path: 'products',
    loadChildren: () => import('./features/products/routes').then(m => m.productRoutes),
    canActivate: [authGuard]
  },
  {
    path: 'admin',
    loadComponent: () => import('./features/admin/dashboard').then(m => m.AdminDashboard),
    canMatch: [() => inject(AuthService).isAdmin()],
    children: [
      {
        path: 'users',
        loadComponent: () => import('./features/admin/users').then(m => m.UserManagement)
      }
    ]
  },
  {
    path: '**',
    loadComponent: () => import('./shared/ui/not-found').then(m => m.NotFound)
  }
];

// features/products/routes.ts
export const productRoutes: Routes = [
  {
    path: '',
    loadComponent: () => import('./list').then(m => m.ProductList),
  },
  {
    path: ':id',
    loadComponent: () => import('./detail').then(m => m.ProductDetail),
    resolve: {
      product: (route: ActivatedRouteSnapshot) => {
        const productService = inject(ProductService);
        return productService.getProduct(route.params['id']);
      }
    }
  }
];

// Guard with signals
export const authGuard = () => {
  const authService = inject(AuthService);
  const router = inject(Router);
  
  return authService.isAuthenticated() 
    ? true 
    : router.createUrlTree(['/login']);
};
```

### HTTP Interceptors with Signals

```typescript
// auth.interceptor.ts
import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const authService = inject(AuthService);
  const token = authService.token();
  
  if (token) {
    req = req.clone({
      headers: req.headers.set('Authorization', `Bearer ${token}`)
    });
  }
  
  return next(req);
};

// error.interceptor.ts
export const errorInterceptor: HttpInterceptorFn = (req, next) => {
  const toastService = inject(ToastService);
  
  return next(req).pipe(
    catchError((error: HttpErrorResponse) => {
      if (error.status === 401) {
        inject(Router).navigate(['/login']);
      } else if (error.status >= 500) {
        toastService.error('Server error occurred');
      }
      
      return throwError(() => error);
    })
  );
};

// retry.interceptor.ts
export const retryInterceptor: HttpInterceptorFn = (req, next) => {
  const maxRetries = 3;
  
  return next(req).pipe(
    retry({
      count: maxRetries,
      delay: (error, retryCount) => {
        if (error.status >= 500) {
          // Exponential backoff: 1s, 2s, 4s
          return timer(Math.pow(2, retryCount - 1) * 1000);
        }
        // Don't retry client errors
        return throwError(() => error);
      }
    })
  );
};
```

## Project Structure

```
src/
\u251c\u2500\u2500 app/
\u2502   \u251c\u2500\u2500 core/
\u2502   \u2502   \u251c\u2500\u2500 services/
\u2502   \u2502   \u2514\u2500\u2500 guards/
\u2502   \u251c\u2500\u2500 features/
\u2502   \u2502   \u251c\u2500\u2500 products/
\u2502   \u2502   \u2502   \u251c\u2500\u2500 list.ts
\u2502   \u2502   \u2502   \u251c\u2500\u2500 detail.ts
\u2502   \u2502   \u2502   \u2514\u2500\u2500 product.service.ts
\u2502   \u2502   \u2514\u2500\u2500 users/
\u2502   \u251c\u2500\u2500 shared/
\u2502   \u2502   \u251c\u2500\u2500 ui/
\u2502   \u2502   \u2514\u2500\u2500 utils/
\u2502   \u2514\u2500\u2500 app.config.ts
```

## Key Advantages Over React

1. **Complete Framework**: Built-in router, forms, HTTP client, testing utilities
2. **TypeScript First**: Superior type safety and IDE support
3. **Consistent Patterns**: Standardized approaches reduce decision fatigue
4. **Enterprise Ready**: 18-month LTS cycles, automated migrations
5. **Performance**: Zoneless mode, incremental hydration, optimal bundle sizes
6. **Developer Experience**: Comprehensive CLI, clear error messages, powerful DevTools

## Resources

- Official Documentation: https://angular.dev
- Migration Guide: https://angular.dev/update-guide
- Angular DevTools: Chrome/Firefox extension
- Community: Discord, Reddit, Stack Overflow

Angular v20 combines enterprise-grade stability with modern performance optimizations, making it an excellent choice for teams building complex, maintainable web applications in 2025.
