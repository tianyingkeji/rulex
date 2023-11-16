// Copyright (C) 2023 wwhai
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package model

var SysDefaultLogo = `data:image/png;base64,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`
