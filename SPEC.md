openapi: 3.0.0
info:
  title: August API
  description: |
    The August API provides programmatic access to legal document intelligence and research workflows.
    
    Non-Obvious Insight: August isn't just a document search tool. It's a legal knowledge graph.
    Every file, folder, and query is a signal about how a legal team structures, retrieves, and
    generates work product — enabling compound insights no single API wrapper could produce.
  version: 1.0.0
  contact:
    name: August Law
    url: https://august.law

servers:
  - url: https://api.august.law
    description: Production

security:
  - BearerAuth: []

paths:
  /api/v1/projects:
    get:
      operationId: listProjects
      summary: List all projects
      description: Returns all projects the authenticated user has access to.
      security:
        - BearerAuth: []
      responses:
        '200':
          description: List of projects
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ProjectList'

  /api/v1/projects/{project_id}/members:
    get:
      operationId: listProjectMembers
      summary: List project members
      description: Returns all members of a specific project.
      security:
        - BearerAuth: []
      parameters:
        - name: project_id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: List of project members
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/MemberList'

  /api/v1/projects/{project_id}/search:
    get:
      operationId: searchProject
      x-pp-resource: project-search
      summary: Search within a project
      description: Full-text search within a specific project's documents and files.
      security:
        - BearerAuth: []
      parameters:
        - name: project_id
          in: path
          required: true
          schema:
            type: string
        - name: q
          in: query
          required: true
          schema:
            type: string
          description: Search query string
        - name: offset
          in: query
          schema:
            type: integer
            default: 0
        - name: limit
          in: query
          schema:
            type: integer
            default: 50
      responses:
        '200':
          description: Search results
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/SearchResults'

  /api/v1/search:
    get:
      operationId: searchGlobal
      x-pp-resource: global-search
      summary: Global search
      description: Full-text search across all accessible projects.
      security:
        - BearerAuth: []
      parameters:
        - name: q
          in: query
          required: true
          schema:
            type: string
        - name: offset
          in: query
          schema:
            type: integer
            default: 0
        - name: limit
          in: query
          schema:
            type: integer
            default: 50
      responses:
        '200':
          description: Search results
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/SearchResults'

  /api/v1/folders/{folder_id}/contents:
    get:
      operationId: getFolderContents
      summary: Get folder contents
      description: Returns the contents (files and subfolders) of a specific folder.
      security:
        - BearerAuth: []
      parameters:
        - name: folder_id
          in: path
          required: true
          schema:
            type: string
        - name: offset
          in: query
          schema:
            type: integer
            default: 0
        - name: limit
          in: query
          schema:
            type: integer
            default: 100
      responses:
        '200':
          description: Folder contents
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/FolderContents'

  /api/v1/folders/{folder_id}/tree:
    get:
      operationId: getFolderTree
      summary: Get folder tree
      description: Returns the full folder tree starting from a specific folder.
      security:
        - BearerAuth: []
      parameters:
        - name: folder_id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Folder tree
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/FolderTree'

  /api/v1/files/download:
    post:
      operationId: getDownloadUrls
      summary: Get presigned download URLs
      description: Generate presigned URLs for downloading one or more files.
      security:
        - BearerAuth: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/DownloadRequest'
      responses:
        '200':
          description: Presigned download URLs
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/DownloadResponse'

  /api/v1/queries:
    post:
      operationId: submitQuery
      summary: Submit a Genius Mode query
      description: |
        Submit a legal research or document analysis query in Genius Mode.
        The query is processed asynchronously; poll status for completion.
      security:
        - BearerAuth: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/QueryRequest'
      responses:
        '200':
          description: Query submitted
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/QuerySubmitResponse'

  /api/v1/queries/{query_id}/status:
    get:
      operationId: getQueryStatus
      summary: Get query status
      description: Poll the processing status of a submitted query.
      security:
        - BearerAuth: []
      parameters:
        - name: query_id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Query status
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/QueryStatus'

  /api/v1/queries/{query_id}/result:
    get:
      operationId: getQueryResult
      summary: Get query result
      description: Retrieve the completed result of a processed query.
      security:
        - BearerAuth: []
      parameters:
        - name: query_id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Query result
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/QueryResult'

  /api/v1/queries/{query_id}/files:
    get:
      operationId: getQueryFiles
      summary: Get generated files
      description: Retrieve files generated by a completed query.
      security:
        - BearerAuth: []
      parameters:
        - name: query_id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Generated files list
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/QueryFiles'

components:
  securitySchemes:
    BearerAuth:
      type: http
      scheme: bearer
      bearerFormat: API Key
      description: August API key (ak_...). Set AUGUST_API_KEY environment variable.

  schemas:
    Project:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        created_at:
          type: string
          format: date-time
        updated_at:
          type: string
          format: date-time

    ProjectList:
      type: object
      properties:
        projects:
          type: array
          items:
            $ref: '#/components/schemas/Project'

    Member:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        email:
          type: string
        role:
          type: string

    MemberList:
      type: object
      properties:
        members:
          type: array
          items:
            $ref: '#/components/schemas/Member'

    SearchResult:
      type: object
      properties:
        id:
          type: string
        type:
          type: string
          enum: [file, folder, document]
        title:
          type: string
        snippet:
          type: string
        project_id:
          type: string
        folder_id:
          type: string
        score:
          type: number

    SearchResults:
      type: object
      properties:
        results:
          type: array
          items:
            $ref: '#/components/schemas/SearchResult'
        total:
          type: integer
        offset:
          type: integer
        limit:
          type: integer

    FolderItem:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        type:
          type: string
          enum: [file, folder]
        size:
          type: integer
        created_at:
          type: string
          format: date-time
        updated_at:
          type: string
          format: date-time

    FolderContents:
      type: object
      properties:
        items:
          type: array
          items:
            $ref: '#/components/schemas/FolderItem'
        offset:
          type: integer
        limit:
          type: integer
        total:
          type: integer

    FolderTreeNode:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        type:
          type: string
          enum: [file, folder]
        children:
          type: array
          items:
            $ref: '#/components/schemas/FolderTreeNode'

    FolderTree:
      type: object
      properties:
        root:
          $ref: '#/components/schemas/FolderTreeNode'

    DownloadRequest:
      type: object
      required:
        - doc_ids
      properties:
        doc_ids:
          type: array
          items:
            type: string
          description: Document IDs to download
        use_pdf:
          type: boolean
          default: false
          description: Download as PDF instead of original format
        ttl:
          type: integer
          default: 3600
          description: URL time-to-live in seconds
        project_id:
          type: string
          description: Optional project scope for the download

    DownloadUrl:
      type: object
      properties:
        doc_id:
          type: string
        url:
          type: string
          format: uri
        expires_at:
          type: string
          format: date-time
        filename:
          type: string

    DownloadResponse:
      type: object
      properties:
        urls:
          type: array
          items:
            $ref: '#/components/schemas/DownloadUrl'

    QueryRequest:
      type: object
      required:
        - query
      properties:
        query:
          type: string
          description: The natural language or structured query
        folder_ids:
          type: array
          items:
            type: string
          description: Folder IDs to scope the query to
        file_ids:
          type: array
          items:
            type: string
          description: File/document IDs to include in the query context
        project_id:
          type: string
          description: Project ID to scope the query to
        mode:
          type: string
          enum: [research, analysis, draft, review]
          default: research
          description: Query processing mode

    QuerySubmitResponse:
      type: object
      properties:
        query_id:
          type: string
        status:
          type: string
        created_at:
          type: string
          format: date-time

    QueryStatus:
      type: object
      properties:
        query_id:
          type: string
        status:
          type: string
          enum: [pending, processing, completed, failed]
        progress:
          type: number
        updated_at:
          type: string
          format: date-time

    QueryResult:
      type: object
      properties:
        query_id:
          type: string
        status:
          type: string
        answer:
          type: string
        sources:
          type: array
          items:
            type: string
        citations:
          type: array
          items:
            $ref: '#/components/schemas/SearchResult'

    QueryFile:
      type: object
      properties:
        id:
          type: string
        name:
          type: string
        type:
          type: string
        size:
          type: integer
        created_at:
          type: string
          format: date-time

    QueryFiles:
      type: object
      properties:
        query_id:
          type: string
        files:
          type: array
          items:
            $ref: '#/components/schemas/QueryFile'
