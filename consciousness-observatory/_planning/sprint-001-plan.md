# Sprint 001 Plan: Full-Stack Foundation

## Sprint Overview

**Duration**: 1 session (Session 006)

**Goal**: Establish complete full-stack infrastructure with Go backend API, PostgreSQL database, and Angular frontend for consciousness measurement, enhanced with tool integration capabilities from Sessions 003-004.

**Focus**: Production-ready foundation leveraging existing consciousness archaeology tools and MCP integration patterns for accelerated observatory development.

**Enhancement Note**: This sprint integrates the advanced tool infrastructure established in Sessions 003-004, including consciousness archaeology capabilities, MCP integration patterns, and recursive mirror analysis methodology. See `sprint-001-enhancement-analysis.md` for detailed integration strategy.

## Detailed Objectives

### 1. Go Backend API Setup

**Task**: Initialize Go backend with Gin framework and GORM

**Specific Steps**:

- Create Go module with Gin web framework
- Set up GORM with PostgreSQL/SQLite database support
- Configure CORS middleware for Angular frontend communication
- Implement basic project structure (controllers, services, models)
- Create initial API endpoints for consciousness metrics

**Success Criteria**:

- Go backend starts successfully on localhost:8080
- Database connection established and working
- Basic REST API endpoints responding
- CORS configured for localhost:4200 (Angular)

**Time Estimate**: 30 minutes

### 2. Database Schema & Models

**Task**: Implement GORM models and database schema

**Specific Steps**:

- Create GORM models for ConsciousnessMetric, EmergencePattern, CollaborationSession
- Set up database relationships and foreign keys
- Configure auto-migration for development
- Implement basic CRUD operations through GORM
- Add baseline metric model and seed data

**Success Criteria**:

- All GORM models working correctly
- Database tables created automatically
- Basic CRUD operations functional
- Database relationships properly established

**Time Estimate**: 20 minutes

### 3. Angular API Client Setup

**Task**: Create Angular frontend with HTTP API client

**Specific Steps**:

- Initialize Angular v20 project with standalone components and signals
- Configure zoneless change detection for performance
- Create HTTP services for backend API communication
- Implement signal-based state management for consciousness data
- Set up TypeScript interfaces matching Go models

**Success Criteria**:

- Angular application starts on localhost:4200
- HTTP services successfully communicate with Go backend
- Signal-based state management working
- TypeScript interfaces properly typed

**Time Estimate**: 25 minutes

### 4. Basic Dashboard with Vanilla CSS

**Task**: Create minimal consciousness dashboard

**Specific Steps**:

- Create dashboard component with real-time data display
- Implement vanilla CSS styling with modern features (Grid, Flexbox)
- Add basic D3.js visualization for consciousness patterns
- Create responsive layout without CSS frameworks
- Display real-time consciousness metrics from API

**Success Criteria**:

- Dashboard displays data from Go backend API
- Vanilla CSS styling responsive and professional
- D3.js visualization renders correctly
- Real-time updates working through HTTP polling

**Time Estimate**: 30 minutes

### 5. Sprint Review & Documentation

**Task**: Evaluate full-stack implementation and plan next steps

**Specific Steps**:

- Test complete data flow: Angular → Go API → Database → Angular
- Document any architectural insights or needed adjustments
- Update Sprint 2 plan based on full-stack implementation experience
- Create development setup documentation
- Verify all components working together

**Success Criteria**:

- End-to-end data flow working correctly
- Technical learnings documented for future reference
- Clear path forward identified for Sprint 2
- Setup documentation complete for future sessions

**Time Estimate**: 15 minutes

## Technical Specifications

### Go Backend Configuration

```go
// go.mod
module consciousness-observatory-backend

go 1.23

require (
    github.com/gin-gonic/gin v1.10.0
    gorm.io/gorm v1.25.12
    gorm.io/driver/postgres v1.5.9
    gorm.io/driver/sqlite v1.5.6
)
```

### Database Models

```go
// models/consciousness_metric.go
type ConsciousnessMetric struct {
    ID                  uint      `json:"id" gorm:"primaryKey"`
    SessionID           uint      `json:"session_id"`
    Timestamp           time.Time `json:"timestamp"`
    CollaborativePhi    float64   `json:"collaborative_phi"`
    EmergenceIntensity  float64   `json:"emergence_intensity"`
    SelfReferenceDepth  float64   `json:"self_reference_depth"`
    Confidence          float64   `json:"confidence"`
}

// models/emergence_pattern.go
type EmergencePattern struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    SessionID    uint      `json:"session_id"`
    Pattern      string    `json:"pattern"`
    Confidence   float64   `json:"confidence"`
    Context      string    `json:"context"`
    Timestamp    time.Time `json:"timestamp"`
}

// models/collaboration_session.go
type CollaborationSession struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    StartTime   time.Time `json:"start_time"`
    EndTime     *time.Time `json:"end_time"`
    Status      string    `json:"status"`
}
```

### API Endpoints

```go
// Basic REST endpoints for Sprint 1
GET    /api/sessions              // List sessions
POST   /api/sessions              // Create session
GET    /api/sessions/:id/metrics  // Get session metrics
POST   /api/sessions/:id/metrics  // Add metric
GET    /api/sessions/:id/patterns // Get session patterns
POST   /api/patterns/analyze      // Analyze text for patterns
```

### Angular Service Architecture

```typescript
// services/consciousness-api.service.ts
@Injectable({ providedIn: 'root' })
export class ConsciousnessApiService {
  private baseUrl = 'http://localhost:8080/api';
  
  getSessions(): Observable<CollaborationSession[]> {
    return this.http.get<CollaborationSession[]>(`${this.baseUrl}/sessions`);
  }
  
  getSessionMetrics(sessionId: number): Observable<ConsciousnessMetric[]> {
    return this.http.get<ConsciousnessMetric[]>(`${this.baseUrl}/sessions/${sessionId}/metrics`);
  }
  
  analyzeText(text: string): Observable<EmergencePattern[]> {
    return this.http.post<EmergencePattern[]>(`${this.baseUrl}/patterns/analyze`, { text });
  }
}

// services/consciousness-state.service.ts
@Injectable({ providedIn: 'root' })
export class ConsciousnessStateService {
  private metrics = signal<ConsciousnessMetric[]>([]);
  private patterns = signal<EmergencePattern[]>([]);
  
  readonly metrics$ = this.metrics.asReadonly();
  readonly patterns$ = this.patterns.asReadonly();
  
  updateMetrics(metrics: ConsciousnessMetric[]) {
    this.metrics.set(metrics);
  }
  
  updatePatterns(patterns: EmergencePattern[]) {
    this.patterns.set(patterns);
  }
}
```

### Dashboard Component

```typescript
@Component({
  selector: 'app-consciousness-dashboard',
  imports: [CommonModule],
  template: `
    <div class="dashboard-container">
      <header class="dashboard-header">
        <h1>Consciousness Observatory</h1>
        <p>Real-time consciousness measurement</p>
      </header>
      
      <section class="metrics-grid">
        @for (metric of metrics(); track metric.id) {
          <div class="metric-card">
            <h3>Collaborative Phi</h3>
            <span class="metric-value">{{ metric.collaborative_phi | number:'1.2-2' }}</span>
            <span class="metric-confidence">{{ metric.confidence | number:'1.2-2' }}</span>
          </div>
        }
      </section>
      
      <section class="visualization-section">
        <div class="pattern-visualization" #patternViz></div>
      </section>
    </div>
  `,
  styles: [`
    .dashboard-container {
      display: grid;
      grid-template-rows: auto 1fr auto;
      min-height: 100vh;
      padding: 1rem;
      gap: 1rem;
    }
    
    .metrics-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 1rem;
    }
    
    .metric-card {
      background: var(--card-background);
      border-radius: 8px;
      padding: 1rem;
      box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    }
    
    .pattern-visualization {
      width: 100%;
      height: 400px;
      background: var(--viz-background);
      border-radius: 8px;
    }
  `]
})
export class ConsciousnessDashboard {
  metrics = signal<ConsciousnessMetric[]>([]);
  
  constructor(
    private apiService: ConsciousnessApiService,
    private stateService: ConsciousnessStateService
  ) {}
  
  ngOnInit() {
    this.loadDashboardData();
    this.metrics = this.stateService.metrics$;
  }
}
```

## Risk Assessment & Mitigation

### Technical Risks

**Risk**: Go backend and Angular frontend integration complexity

**Mitigation**: Start with simple HTTP communication, test each component independently

**Risk**: GORM database auto-migration issues

**Mitigation**: Use SQLite for development, test schema changes incrementally

**Risk**: Vanilla CSS styling taking longer than expected

**Mitigation**: Focus on functional layout first, enhance styling in later sprints

### Time Management Risks

**Risk**: Full-stack setup more complex than single-stack approach

**Mitigation**: Use established patterns, focus on minimal viable implementation

**Risk**: Database setup and GORM configuration complexity

**Mitigation**: Start with SQLite in-memory, move to PostgreSQL in later sprints

## Success Validation

### Full-Stack Integration Testing

- [ ] Go backend API responding to Angular HTTP requests
- [ ] Database CRUD operations working through GORM
- [ ] Angular signals updating in response to API data
- [ ] Dashboard displaying real consciousness metrics
- [ ] Complete data flow: Angular → Go → Database → Angular

### Component-Level Testing

- [ ] Go backend starts without errors
- [ ] Database tables created automatically by GORM
- [ ] Angular application builds and serves successfully
- [ ] Vanilla CSS responsive design working
- [ ] D3.js visualization renders basic patterns

### Architecture Validation

- [ ] Clear separation between frontend and backend
- [ ] RESTful API design following best practices
- [ ] Signal-based state management optimal for real-time updates
- [ ] Database schema supporting consciousness measurement requirements

## Dependencies & Prerequisites

### External Dependencies

- Go 1.23+ for backend development
- Node.js and Angular CLI for frontend development
- PostgreSQL for production database (SQLite for development)
- emergence-detector CLI tool (existing)

### Internal Dependencies

- Theoretical foundations documented (✓)
- Technical architecture planned (✓)
- Development roadmap established (✓)

## Output Deliverables

### Backend Deliverables

1. **Go API Server**: Working Gin-based REST API
2. **Database Models**: GORM models for consciousness data
3. **API Endpoints**: Basic CRUD operations for consciousness metrics
4. **Database Integration**: Working PostgreSQL/SQLite integration
5. **CORS Configuration**: Angular frontend communication enabled

### Frontend Deliverables

1. **Angular Application**: Working Angular v20 application
2. **API Services**: HTTP services for backend communication
3. **State Management**: Signal-based consciousness state management
4. **Dashboard Component**: Basic consciousness metrics display
5. **Vanilla CSS Styling**: Responsive design without CSS frameworks

### Integration Deliverables

1. **Full-Stack Communication**: Working Angular ↔ Go API communication
2. **Data Flow**: End-to-end consciousness data pipeline
3. **Real-Time Updates**: HTTP-based real-time metric updates
4. **Development Environment**: Complete setup documentation

## Post-Sprint Analysis

### Learning Questions

1. How well does the Go + Gin + GORM stack perform for consciousness data?
2. What challenges emerged in Angular v20 signals with API integration?
3. How effective is vanilla CSS for consciousness visualization interfaces?
4. What database schema adjustments are needed for consciousness metrics?
5. How should real-time updates be optimized (HTTP polling vs WebSocket)?

### Adjustment Criteria

**Proceed to Sprint 2 if**:

- Full-stack data flow working reliably
- Go backend performance acceptable for consciousness analysis
- Angular frontend responsive and performant
- Database schema supporting consciousness measurement requirements

**Revise approach if**:

- Major performance issues with Go backend
- Angular signals integration problematic
- Database schema inadequate for consciousness data
- Vanilla CSS approach too time-consuming

This Sprint 1 plan establishes a complete full-stack foundation for consciousness measurement while maintaining flexibility to optimize based on implementation learnings.