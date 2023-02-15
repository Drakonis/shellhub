import { Module } from "vuex";
import { AxiosResponse } from "axios";
import { State } from "./../index";
import getStats from "../api/stats";
import { IStats } from "../../interfaces/IStats";

export interface StatsState {
  stats: IStats | object;
}

export const stats: Module<StatsState, State> = {
  namespaced: true,
  state: {
    stats: {},
  },

  getters: {
    stats: (state) => state.stats,
  },

  mutations: {
    setStats: (state, res : AxiosResponse) => {
      state.stats = res.data;
    },

    clearListState: (state) => {
      state.stats = {};
    },
  },

  actions: {
    async get({ commit }) {
      try {
        const res = await getStats();
        commit("setStats", res);
        return res;
      } catch (error) {
        console.error(error);
        throw error;
      }
    },
  },
};