import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { map } from "rxjs/operators";
import { HttpClient } from "@angular/common/http";
import { PlayerProfile } from "../models/profile";


@Injectable({ providedIn: "root" })
export class SearchService {
  constructor(private readonly http: HttpClient) {}

  searchProfile(nickname: string, patch: string): Observable<PlayerProfile> {
    return this.http
      .post<{ profile: PlayerProfile }>(`/player`, {
        nickname: nickname,
        patch: patch,
        side: ["dire","radiant"],
        wards: ["sentry", "observer"],
        time: [0, 1000],
      })
      .pipe(map((data) => data.profile));
  }
}
