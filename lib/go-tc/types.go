package tc

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// TypesResponse ...
type TypesResponse struct {
	Response []Type `json:"response"`
}

// Type contains information about a given Type in Traffic Ops.
type Type struct {
	ID          int       `json:"id"`
	LastUpdated TimeNoMod `json:"lastUpdated"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UseInTable  string    `json:"useInTable"`
}

// TypeNullable contains information about a given Type in Traffic Ops.
type TypeNullable struct {
	ID          *int       `json:"id"`
	LastUpdated *TimeNoMod `json:"lastUpdated" db:"last_updated"`
	Name        *string    `json:"name" db:"name"`
	Description *string    `json:"description" db:"description"`
	UseInTable  *string    `json:"useInTable" db:"use_in_table"`
}
